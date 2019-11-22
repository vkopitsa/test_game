// package: messages
// file: command.proto

import * as jspb from "google-protobuf";

export class Command extends jspb.Message {
  getTime(): number;
  setTime(value: number): void;

  getYv(): number;
  setYv(value: number): void;

  getXv(): number;
  setXv(value: number): void;

  getCmd(): Command.CmdMap[keyof Command.CmdMap];
  setCmd(value: Command.CmdMap[keyof Command.CmdMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Command.AsObject;
  static toObject(includeInstance: boolean, msg: Command): Command.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Command, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Command;
  static deserializeBinaryFromReader(message: Command, reader: jspb.BinaryReader): Command;
}

export namespace Command {
  export type AsObject = {
    time: number,
    yv: number,
    xv: number,
    cmd: Command.CmdMap[keyof Command.CmdMap],
  }

  export interface CmdMap {
    NONE: 0;
    SHOT: 1;
  }

  export const Cmd: CmdMap;
}

