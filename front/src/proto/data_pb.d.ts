// package: messages
// file: data.proto

import * as jspb from "google-protobuf";

export class Data extends jspb.Message {
  getY(): number;
  setY(value: number): void;

  getX(): number;
  setX(value: number): void;

  getYv(): number;
  setYv(value: number): void;

  getXv(): number;
  setXv(value: number): void;

  getHeight(): number;
  setHeight(value: number): void;

  getWidth(): number;
  setWidth(value: number): void;

  getColor(): string;
  setColor(value: string): void;

  getTime(): number;
  setTime(value: number): void;

  getDelta(): number;
  setDelta(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Data.AsObject;
  static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Data;
  static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
}

export namespace Data {
  export type AsObject = {
    y: number,
    x: number,
    yv: number,
    xv: number,
    height: number,
    width: number,
    color: string,
    time: number,
    delta: number,
  }
}

