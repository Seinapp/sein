export interface IClient {
  close(): void;
  updateToken(token: string): void;
}
