// package: auth
// file: auth.proto

/* tslint:disable */

import * as grpc from "grpc";
import * as auth_pb from "./auth_pb";

interface IAuthService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    signIn: IAuthService_ISignIn;
    signOut: IAuthService_ISignOut;
    keepAlive: IAuthService_IKeepAlive;
}

interface IAuthService_ISignIn extends grpc.MethodDefinition<auth_pb.SignInRequest, auth_pb.SignInResponse> {
    path: string; // "/auth.Auth/SignIn"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<auth_pb.SignInRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.SignInRequest>;
    responseSerialize: grpc.serialize<auth_pb.SignInResponse>;
    responseDeserialize: grpc.deserialize<auth_pb.SignInResponse>;
}
interface IAuthService_ISignOut extends grpc.MethodDefinition<auth_pb.SignOutRequest, auth_pb.SignOutResponse> {
    path: string; // "/auth.Auth/SignOut"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<auth_pb.SignOutRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.SignOutRequest>;
    responseSerialize: grpc.serialize<auth_pb.SignOutResponse>;
    responseDeserialize: grpc.deserialize<auth_pb.SignOutResponse>;
}
interface IAuthService_IKeepAlive extends grpc.MethodDefinition<auth_pb.KeepAliveRequest, auth_pb.KeepAliveResponse> {
    path: string; // "/auth.Auth/KeepAlive"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<auth_pb.KeepAliveRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.KeepAliveRequest>;
    responseSerialize: grpc.serialize<auth_pb.KeepAliveResponse>;
    responseDeserialize: grpc.deserialize<auth_pb.KeepAliveResponse>;
}

export const AuthService: IAuthService;

export interface IAuthServer {
    signIn: grpc.handleUnaryCall<auth_pb.SignInRequest, auth_pb.SignInResponse>;
    signOut: grpc.handleUnaryCall<auth_pb.SignOutRequest, auth_pb.SignOutResponse>;
    keepAlive: grpc.handleUnaryCall<auth_pb.KeepAliveRequest, auth_pb.KeepAliveResponse>;
}

export interface IAuthClient {
    signIn(request: auth_pb.SignInRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    signIn(request: auth_pb.SignInRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    signIn(request: auth_pb.SignInRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    signOut(request: auth_pb.SignOutRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.SignOutResponse) => void): grpc.ClientUnaryCall;
    signOut(request: auth_pb.SignOutRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.SignOutResponse) => void): grpc.ClientUnaryCall;
    signOut(request: auth_pb.SignOutRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.SignOutResponse) => void): grpc.ClientUnaryCall;
    keepAlive(request: auth_pb.KeepAliveRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.KeepAliveResponse) => void): grpc.ClientUnaryCall;
    keepAlive(request: auth_pb.KeepAliveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.KeepAliveResponse) => void): grpc.ClientUnaryCall;
    keepAlive(request: auth_pb.KeepAliveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.KeepAliveResponse) => void): grpc.ClientUnaryCall;
}

export class AuthClient extends grpc.Client implements IAuthClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public signIn(request: auth_pb.SignInRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    public signIn(request: auth_pb.SignInRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    public signIn(request: auth_pb.SignInRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.SignInResponse) => void): grpc.ClientUnaryCall;
    public signOut(request: auth_pb.SignOutRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.SignOutResponse) => void): grpc.ClientUnaryCall;
    public signOut(request: auth_pb.SignOutRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.SignOutResponse) => void): grpc.ClientUnaryCall;
    public signOut(request: auth_pb.SignOutRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.SignOutResponse) => void): grpc.ClientUnaryCall;
    public keepAlive(request: auth_pb.KeepAliveRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.KeepAliveResponse) => void): grpc.ClientUnaryCall;
    public keepAlive(request: auth_pb.KeepAliveRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.KeepAliveResponse) => void): grpc.ClientUnaryCall;
    public keepAlive(request: auth_pb.KeepAliveRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.KeepAliveResponse) => void): grpc.ClientUnaryCall;
}
