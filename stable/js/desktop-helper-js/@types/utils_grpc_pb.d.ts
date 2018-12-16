// package: utils
// file: utils.proto

/* tslint:disable */

import * as grpc from "grpc";
import * as utils_pb from "./utils_pb";

interface IUtilsService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    ping: IUtilsService_IPing;
}

interface IUtilsService_IPing extends grpc.MethodDefinition<utils_pb.PingRequest, utils_pb.PingResponse> {
    path: string; // "/utils.Utils/Ping"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestSerialize: grpc.serialize<utils_pb.PingRequest>;
    requestDeserialize: grpc.deserialize<utils_pb.PingRequest>;
    responseSerialize: grpc.serialize<utils_pb.PingResponse>;
    responseDeserialize: grpc.deserialize<utils_pb.PingResponse>;
}

export const UtilsService: IUtilsService;

export interface IUtilsServer {
    ping: grpc.handleUnaryCall<utils_pb.PingRequest, utils_pb.PingResponse>;
}

export interface IUtilsClient {
    ping(request: utils_pb.PingRequest, callback: (error: grpc.ServiceError | null, response: utils_pb.PingResponse) => void): grpc.ClientUnaryCall;
    ping(request: utils_pb.PingRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: utils_pb.PingResponse) => void): grpc.ClientUnaryCall;
    ping(request: utils_pb.PingRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: utils_pb.PingResponse) => void): grpc.ClientUnaryCall;
}

export class UtilsClient extends grpc.Client implements IUtilsClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public ping(request: utils_pb.PingRequest, callback: (error: grpc.ServiceError | null, response: utils_pb.PingResponse) => void): grpc.ClientUnaryCall;
    public ping(request: utils_pb.PingRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: utils_pb.PingResponse) => void): grpc.ClientUnaryCall;
    public ping(request: utils_pb.PingRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: utils_pb.PingResponse) => void): grpc.ClientUnaryCall;
}
