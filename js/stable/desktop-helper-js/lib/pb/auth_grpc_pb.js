// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var auth_pb = require('./auth_pb.js');

function serialize_auth_KeepAliveRequest(arg) {
  if (!(arg instanceof auth_pb.KeepAliveRequest)) {
    throw new Error('Expected argument of type auth.KeepAliveRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_KeepAliveRequest(buffer_arg) {
  return auth_pb.KeepAliveRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_KeepAliveResponse(arg) {
  if (!(arg instanceof auth_pb.KeepAliveResponse)) {
    throw new Error('Expected argument of type auth.KeepAliveResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_KeepAliveResponse(buffer_arg) {
  return auth_pb.KeepAliveResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_SignInRequest(arg) {
  if (!(arg instanceof auth_pb.SignInRequest)) {
    throw new Error('Expected argument of type auth.SignInRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_SignInRequest(buffer_arg) {
  return auth_pb.SignInRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_SignInResponse(arg) {
  if (!(arg instanceof auth_pb.SignInResponse)) {
    throw new Error('Expected argument of type auth.SignInResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_SignInResponse(buffer_arg) {
  return auth_pb.SignInResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_SignOutRequest(arg) {
  if (!(arg instanceof auth_pb.SignOutRequest)) {
    throw new Error('Expected argument of type auth.SignOutRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_SignOutRequest(buffer_arg) {
  return auth_pb.SignOutRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_SignOutResponse(arg) {
  if (!(arg instanceof auth_pb.SignOutResponse)) {
    throw new Error('Expected argument of type auth.SignOutResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_auth_SignOutResponse(buffer_arg) {
  return auth_pb.SignOutResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Auth represents the authentication server
var AuthService = exports.AuthService = {
  signIn: {
    path: '/auth.Auth/SignIn',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.SignInRequest,
    responseType: auth_pb.SignInResponse,
    requestSerialize: serialize_auth_SignInRequest,
    requestDeserialize: deserialize_auth_SignInRequest,
    responseSerialize: serialize_auth_SignInResponse,
    responseDeserialize: deserialize_auth_SignInResponse,
  },
  signOut: {
    path: '/auth.Auth/SignOut',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.SignOutRequest,
    responseType: auth_pb.SignOutResponse,
    requestSerialize: serialize_auth_SignOutRequest,
    requestDeserialize: deserialize_auth_SignOutRequest,
    responseSerialize: serialize_auth_SignOutResponse,
    responseDeserialize: deserialize_auth_SignOutResponse,
  },
  keepAlive: {
    path: '/auth.Auth/KeepAlive',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.KeepAliveRequest,
    responseType: auth_pb.KeepAliveResponse,
    requestSerialize: serialize_auth_KeepAliveRequest,
    requestDeserialize: deserialize_auth_KeepAliveRequest,
    responseSerialize: serialize_auth_KeepAliveResponse,
    responseDeserialize: deserialize_auth_KeepAliveResponse,
  },
};

exports.AuthClient = grpc.makeGenericClientConstructor(AuthService);
