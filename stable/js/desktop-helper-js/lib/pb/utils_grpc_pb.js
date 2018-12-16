// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var utils_pb = require('./utils_pb.js');

function serialize_utils_PingRequest(arg) {
  if (!(arg instanceof utils_pb.PingRequest)) {
    throw new Error('Expected argument of type utils.PingRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_utils_PingRequest(buffer_arg) {
  return utils_pb.PingRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_utils_PingResponse(arg) {
  if (!(arg instanceof utils_pb.PingResponse)) {
    throw new Error('Expected argument of type utils.PingResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_utils_PingResponse(buffer_arg) {
  return utils_pb.PingResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Utils represents the a server used to do utilities tasks
var UtilsService = exports.UtilsService = {
  ping: {
    path: '/utils.Utils/Ping',
    requestStream: false,
    responseStream: false,
    requestType: utils_pb.PingRequest,
    responseType: utils_pb.PingResponse,
    requestSerialize: serialize_utils_PingRequest,
    requestDeserialize: deserialize_utils_PingRequest,
    responseSerialize: serialize_utils_PingResponse,
    responseDeserialize: deserialize_utils_PingResponse,
  },
};

exports.UtilsClient = grpc.makeGenericClientConstructor(UtilsService);
