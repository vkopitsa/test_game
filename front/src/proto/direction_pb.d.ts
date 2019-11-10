// package: messages
// file: direction.proto

import * as jspb from "google-protobuf";

export class Direction extends jspb.Message {
  getType(): Direction.TypeMap[keyof Direction.TypeMap];
  setType(value: Direction.TypeMap[keyof Direction.TypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Direction.AsObject;
  static toObject(includeInstance: boolean, msg: Direction): Direction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Direction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Direction;
  static deserializeBinaryFromReader(message: Direction, reader: jspb.BinaryReader): Direction;
}

export namespace Direction {
  export type AsObject = {
    type: Direction.TypeMap[keyof Direction.TypeMap],
  }

  export interface TypeMap {
    STOP: 0;
    UP: 1;
    DOWN: 2;
    LEFT: 3;
    RIGHT: 4;
  }

  export const Type: TypeMap;
}

