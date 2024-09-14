import * as jspb from 'google-protobuf';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"

export class Module extends jspb.Message {
  getId(): number;
  setId(value: number): Module;

  getName(): string;
  setName(value: string): Module;

  getDefaultSystem(): boolean;
  setDefaultSystem(value: boolean): Module;

  getInstalled(): boolean;
  setInstalled(value: boolean): Module;

  getStarted(): boolean;
  setStarted(value: boolean): Module;

  getEnabled(): boolean;
  setEnabled(value: boolean): Module;

  getTool(): string;
  setTool(value: string): Module;

  getToolVersion(): string;
  setToolVersion(value: string): Module;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Module;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Module;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Module;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Module;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Module.AsObject;
  static toObject(includeInstance: boolean, msg: Module): Module.AsObject;
  static serializeBinaryToWriter(
    message: Module,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): Module;
  static deserializeBinaryFromReader(
    message: Module,
    reader: jspb.BinaryReader
  ): Module;
}

export namespace Module {
  export type AsObject = {
    id: number;
    name: string;
    defaultSystem: boolean;
    installed: boolean;
    started: boolean;
    enabled: boolean;
    tool: string;
    toolVersion: string;
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject;
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject;
  };
}

export class GetModulesResponse extends jspb.Message {
  getResultList(): Array<Module>;
  setResultList(value: Array<Module>): GetModulesResponse;
  clearResultList(): GetModulesResponse;
  addResult(value?: Module, index?: number): Module;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetModulesResponse.AsObject;
  static toObject(
    includeInstance: boolean,
    msg: GetModulesResponse
  ): GetModulesResponse.AsObject;
  static serializeBinaryToWriter(
    message: GetModulesResponse,
    writer: jspb.BinaryWriter
  ): void;
  static deserializeBinary(bytes: Uint8Array): GetModulesResponse;
  static deserializeBinaryFromReader(
    message: GetModulesResponse,
    reader: jspb.BinaryReader
  ): GetModulesResponse;
}

export namespace GetModulesResponse {
  export type AsObject = {
    resultList: Array<Module.AsObject>;
  };
}
