import { spawn } from 'child_process';
import getPort from 'get-port';
import { seinHelperPath } from 'sein-helper';
import util from 'util';
import { HelperError } from './lib/errors';
import { UtilsClient, AuthClient, IClient } from './lib/clients';
import { credentials } from 'grpc';

const exec = util.promisify(spawn);

interface IExec {
  stdout: string;
  stderr: string;
}

interface ICLIstderr {
  error?: string;
  warning?: string;
  info?: string;
  at: string;
}

enum ClientName {
  Auth = 1,
  Utils
}

class SeinServer {
  clients: IClient[] = [];

  private isRunning = false;
  private port = 0;
  private env: string;

  // Important: masterPassword should be cleared when the user
  // stops the server/closes the app
  private masterPassword = '';

  // for security reasons server will close the connection if we don't
  // call keepAlive() before the deadline provided either during the last
  // signIn or during the last keepAlive() call
  private keepAliveWatcherId: NodeJS.Timeout | null = null;
  private authToken = '';

  constructor(env: string = 'production') {
    this.env = env;
  }

  async utilsClient(): Promise<UtilsClient | null> {
    if (!this.isRunning) {
      return null;
    }

    const client = this.clients[ClientName.Utils];
    if (!client) {
      return null;
    }
    return client as UtilsClient;
  }

  // authClient is private to let the class deal with auth itself
  private async authClient(): Promise<AuthClient | null> {
    if (!this.isRunning) {
      await this.start();
    }

    const client = this.clients[ClientName.Auth];
    if (!client) {
      return null;
    }
    return client as AuthClient;
  }

  // signIn signs the client in
  // throws HelperError in case of connection error
  // throws grpc.ServiceError in case of message error
  async signIn(masterPassword: string) {
    const authClient = await this.authClient();
    if (authClient == null) {
      // TODO(melvin): deal with this (maybe authClient() should throw
      // an error)
      return;
    }

    const resp = await authClient.signIn(masterPassword);
    this.masterPassword = masterPassword;
    this.authToken = resp.getClientToken();
    this.clients.forEach(c => {
      c.updateToken(this.authToken);
    });
    this.watchKeepAlive(resp.getAliveUntil());
  }

  // watchKeepAlive creates a listener that makes sure the server doesn't
  // times out
  private watchKeepAlive(triggersAt: number) {
    if (this.keepAliveWatcherId) {
      clearTimeout(this.keepAliveWatcherId);
    }

    // TODO(melvin): makes sure triggersAt is smaller than utcTime
    const utcTime = new Date().getTime();
    const tenSeconds = 10 * 1000;

    this.keepAliveWatcherId = setTimeout(() => {
      this.keepAlive();
    }, utcTime - triggersAt - tenSeconds);
  }

  // keepAlive sends a keepAlive request to the server, making sure
  // the server doesn't stop
  async keepAlive() {
    const authClient = await this.authClient();
    if (authClient == null) {
      // TODO(melvin): deal with this (maybe authClient() should throw
      // an error), but it needs to be logged for sure, somehow
      return;
    }

    // TODO(melvin): handle errors
    const resp = await authClient.keepAlive();
    this.watchKeepAlive(resp.getAliveUntil());
  }

  // monitor pings the gRPC server to check when it dies
  // TODO(melvin): implement
  async monitor() {}

  // start starts the gRPC server on any port available
  // throws HelperError in case of error
  async start() {
    if (this.isRunning) {
      return;
    }
    this.resetData();

    this.port = await getPort();
    // TODO(melvin): check for the port already in use error and retry using
    // another port
    this.cli([
      'start',
      '--json',
      '-d',
      '--env',
      this.env,
      '--port',
      this.port.toString()
    ]);

    this.isRunning = true;
    this.setClients();
  }

  // start stops the gRPC server for the specified env
  // throws HelperError in case of error
  async stop() {
    this.resetData();
    await this.cli(['stop', '--json', '-d', '--env', this.env]);
  }

  // resetData closes all the active connections, and makes this class
  // ready to start a new one
  private resetData() {
    if (this.keepAliveWatcherId) {
      clearTimeout(this.keepAliveWatcherId);
    }
    this.keepAliveWatcherId = null;

    this.masterPassword = '';
    this.authToken = '';
    this.isRunning = false;
    this.clients.forEach(c => {
      c.close();
    });
    this.clients = [];
  }

  // cli runs a Sein Helper cli command
  // throws HelperError in case of error
  async cli(args: ReadonlyArray<string>) {
    const response = (await exec(seinHelperPath, args, undefined)) as IExec;

    // stderr can contain warnings, infos, and errors. Each line contains
    // a valid json object and ends with a "\n"
    // We don't care about anything else than errors.
    // TODO(melvin): It could be useful to log all the messages
    if (response.stderr) {
      const messages = response.stderr.split('\n');
      messages.forEach((msg: string) => {
        let msgObject: ICLIstderr;

        // the message should be a JSON
        try {
          msgObject = JSON.parse(msg);
        } catch (e) {
          throw new HelperError(msg);
        }

        // if we got an actual error then something went wrong
        if (msgObject.error) {
          throw new HelperError(msgObject.error);
        }
      });
    }
  }

  address(): string {
    return `127.0.0.1:${this.port}`;
  }

  // setClients sets the clients for all the proto file
  private setClients() {
    const addr = this.address();
    const creds = credentials.createInsecure();

    this.clients[ClientName.Utils] = new UtilsClient(addr, creds);
    this.clients[ClientName.Auth] = new AuthClient(addr, creds);
  }
}
