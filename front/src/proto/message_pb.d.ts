// package: messages
// file: message.proto

import * as jspb from "google-protobuf";

export class Message extends jspb.Message {
  getType(): Message.TypeMap[keyof Message.TypeMap];
  setType(value: Message.TypeMap[keyof Message.TypeMap]): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getPlayerid(): number;
  setPlayerid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    type: Message.TypeMap[keyof Message.TypeMap],
    data: Uint8Array | string,
    playerid: number,
  }

  export interface TypeMap {
    UNKNOWN: 0;
    HELLO: 1;
    AUTH: 2;
    JOIN: 3;
    QUIT: 4;
    DATA: 5;
    JOINED: 6;
    DIRECTION: 7;
    COMMAND: 8;
  }

  export const Type: TypeMap;
}

