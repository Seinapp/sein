import { IClient } from './client';
import { UtilsClient as GRPCClient } from '../@types/utils_grpc_pb';
import { ChannelCredentials, ServiceError } from 'grpc';
import { PingRequest, PingResponse } from '../@types/utils_pb';

export class UtilsClient implements IClient {
  private client: GRPCClient;

  constructor(
    address: string,
    credentials: ChannelCredentials,
    options?: object | undefined
  ) {
    this.client = new GRPCClient(address, credentials, options);
  }

  ping(): Promise<PingResponse> {
    return new Promise<PingResponse>(
      (
        resolve: (value?: PingResponse | PromiseLike<PingResponse>) => void,
        reject: (reason?: any) => void
      ) => {
        this.client.ping(
          new PingRequest(),
          (err: ServiceError | null, response: PingResponse) => {
            if (err) {
              return reject(err);
            }
            return resolve(response);
          }
        );
      }
    );
  }

  updateToken(token: string) {}

  close(): void {
    this.client.close();
  }
}
