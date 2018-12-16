import { IClient } from './client';
import { AuthClient as GRPCClient } from '../@types/auth_grpc_pb';
import { ChannelCredentials } from 'grpc';
import {
  KeepAliveRequest,
  KeepAliveResponse,
  SignInResponse,
  SignInRequest
} from '../@types/auth_pb';
import { AuthError } from './errors';

export class AuthClient implements IClient {
  private client: GRPCClient;
  private authToken?: string;

  constructor(
    address: string,
    credentials: ChannelCredentials,
    options?: object | undefined
  ) {
    this.client = new GRPCClient(address, credentials, options);
  }

  keepAlive(): Promise<KeepAliveResponse> {
    return new Promise<KeepAliveResponse>((resolve, reject) => {
      if (!this.authToken) {
        return reject(
          new AuthError('you need to sign in before making this request')
        );
      }

      const req = new KeepAliveRequest();
      req.setClientToken(this.authToken);

      this.client.keepAlive(req, (err, response) => {
        if (err) {
          return reject(err);
        }
        return resolve(response);
      });
    });
  }

  signIn(masterPassword: string): Promise<SignInResponse> {
    return new Promise<SignInResponse>((resolve, reject) => {
      const req = new SignInRequest();
      req.setMasterPassword(masterPassword);

      this.client.signIn(req, (err, response) => {
        if (err) {
          return reject(err);
        }
        return resolve(response);
      });
    });
  }

  updateToken(token: string) {
    this.authToken = token;
  }

  close(): void {
    this.client.close();
  }
}
