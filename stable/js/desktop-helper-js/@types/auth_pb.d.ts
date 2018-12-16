// package: auth
// file: auth.proto

/* tslint:disable */

import * as jspb from "google-protobuf";

export class SignInRequest extends jspb.Message { 
    getMasterPassword(): string;
    setMasterPassword(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignInRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SignInRequest): SignInRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SignInRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SignInRequest;
    static deserializeBinaryFromReader(message: SignInRequest, reader: jspb.BinaryReader): SignInRequest;
}

export namespace SignInRequest {
    export type AsObject = {
        masterPassword: string,
    }
}

export class SignInResponse extends jspb.Message { 
    getClientToken(): string;
    setClientToken(value: string): void;

    getAliveUntil(): number;
    setAliveUntil(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignInResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SignInResponse): SignInResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SignInResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SignInResponse;
    static deserializeBinaryFromReader(message: SignInResponse, reader: jspb.BinaryReader): SignInResponse;
}

export namespace SignInResponse {
    export type AsObject = {
        clientToken: string,
        aliveUntil: number,
    }
}

export class SignOutRequest extends jspb.Message { 
    getClientToken(): string;
    setClientToken(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignOutRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SignOutRequest): SignOutRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SignOutRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SignOutRequest;
    static deserializeBinaryFromReader(message: SignOutRequest, reader: jspb.BinaryReader): SignOutRequest;
}

export namespace SignOutRequest {
    export type AsObject = {
        clientToken: string,
    }
}

export class SignOutResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SignOutResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SignOutResponse): SignOutResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SignOutResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SignOutResponse;
    static deserializeBinaryFromReader(message: SignOutResponse, reader: jspb.BinaryReader): SignOutResponse;
}

export namespace SignOutResponse {
    export type AsObject = {
    }
}

export class KeepAliveRequest extends jspb.Message { 
    getClientToken(): string;
    setClientToken(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): KeepAliveRequest.AsObject;
    static toObject(includeInstance: boolean, msg: KeepAliveRequest): KeepAliveRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: KeepAliveRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): KeepAliveRequest;
    static deserializeBinaryFromReader(message: KeepAliveRequest, reader: jspb.BinaryReader): KeepAliveRequest;
}

export namespace KeepAliveRequest {
    export type AsObject = {
        clientToken: string,
    }
}

export class KeepAliveResponse extends jspb.Message { 
    getAliveUntil(): number;
    setAliveUntil(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): KeepAliveResponse.AsObject;
    static toObject(includeInstance: boolean, msg: KeepAliveResponse): KeepAliveResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: KeepAliveResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): KeepAliveResponse;
    static deserializeBinaryFromReader(message: KeepAliveResponse, reader: jspb.BinaryReader): KeepAliveResponse;
}

export namespace KeepAliveResponse {
    export type AsObject = {
        aliveUntil: number,
    }
}
