import { IСommunication } from "./communication";
import { Message } from "./proto/message_pb";
import { Command } from "./proto/command_pb";
import { Direction as PDirection } from "./proto/direction_pb";
import { Player } from './player';
import { Map as GameMap } from './map';
import { Camera } from "./camera";
import { Data } from './proto/data_pb';
import { Info } from './proto/info_pb';
import { GameInfo } from './gameInfo';
import { Control } from './control';

export class Point {constructor(public x = 0, public y = 0) {}};

export enum Direction {
  Stop = 0,
  Up,
  Down,
  Left,
  Right,
  None,
}

export class Game {
  private readonly ctx: CanvasRenderingContext2D; // HTML Canvas's 2D context
  private canvasWidth: number = 0; // width of the canvas
  private canvasHeight: number = 0; // height of the canvas
  private lastTime: number = 0;
  // private gameTime: number = 0;
  private playDirection: Direction = Direction.Stop;
  private previousPlayDirection: Direction = Direction.Stop;
  private playerId: number = 0;
  private readonly players: Map<number,Player> = new Map<number,Player>();

  private map: GameMap;
  private camera: Camera;

  private control: Control;

  private infoEl: GameInfo = new GameInfo();

  private serverDelta: number = 0.0;

  private playerNearValue: number = 1500;


  constructor(canvas: HTMLCanvasElement, private communication: IСommunication) {
    this.ctx = canvas.getContext('2d')!;
    this.resizeCanvasToDisplaySize(this.ctx.canvas);

    this.control = new Control(this.ctx.canvas.width, this.ctx.canvas.height);

    const width = 5000;
    const height = 3000;

    this.map = new GameMap(width, height);
    this.map.generate();

    // Set the right viewport size for the camera
    const vWidth = Math.min(width, canvas.width);
    const vHeight = Math.min(height, canvas.height);

    // Setup the camera
    this.camera = new Camera(0, 0, vWidth, vHeight, width, height);

    // Setup the control
    const that = this;
    window.addEventListener('keydown', function(event) {
      const playDirection = that.control.getDirectionBykeyCode(event.keyCode)
      if (playDirection !== Direction.None) {
        that.playDirection = playDirection
      }
    }, false);

    const elemLeft = canvas.offsetLeft, elemTop = canvas.offsetTop;
    canvas.addEventListener('click', function(event) {
      var x = event.pageX - elemLeft,
        y = event.pageY - elemTop;

        const playDirection = that.control.getDirectionByPoint(new Point(x, y));
        if (playDirection !== Direction.None) {
          that.playDirection = playDirection
        }
    }, false);

    // Processing message
    this.communication.onMesssage((msg: Message) => {
      if (msg.getType() === Message.Type.JOINED && !this.players.has(msg.getPlayerid())) {
        const player = new Player(msg.getPlayerid(), 50, 50);
        player.setDirection(Direction.None)
        
        this.players.set(msg.getPlayerid(), player)
        this.playerId = msg.getPlayerid();

        this.camera.follow(player, vWidth / 2, vHeight / 2);
    } else if (msg.getType() === Message.Type.INFO) {
      const info = Info.deserializeBinary(msg.getData_asU8());
      this.infoEl.setInfo(info)
    } else if (msg.getType() === Message.Type.QUIT) {
      if (this.players.has(msg.getPlayerid())) {
        this.players.delete(msg.getPlayerid())
      }
    } else if (msg.getType() === Message.Type.DATA) {
        const data = Data.deserializeBinary(msg.getData_asU8());

        // set server delta need to interpolation
        this.serverDelta = data.getDelta()

        let player: Player;
        if (this.players.has(msg.getPlayerid())) {
          player = this.players.get(msg.getPlayerid())!
          player.setColor(data.getColor())
        } else {
          player = new Player(msg.getPlayerid(), data.getX(), data.getY(), 200, 50, data.getColor());
          player.setDirection(Direction.None)
          this.players.set(msg.getPlayerid(), player);
        }

        player.applyData(data);
        // if (msg.getPlayerid() === this.playerId) {
        //   player.applyData(data);
        // } else {
        //   player.addData(data);
        // }

        // // reconciliation
        // if (msg.getPlayerid() === this.playerId) {
        //   player.reconciliation(data);
        // } else {
        //   // ineed to interpolation
        //   player.addData(data);
        // }
      }

    });
  }

  public start() {
    var now = Date.now();
    var dt = (now - this.lastTime) / 1000.0;

    this.update(dt);
    this.render();

    this.lastTime = now;
    this.requestAnimationFrame(() => this.start());
  }

  private update(dt: number) {
    // this.gameTime += dt;

    this.handleInput(dt);
    this.updateEntities(dt);
  }

  private handleInput(dt: number) {

    // pass if the same direction
    if (this.playDirection === Direction.None) {
      return
    }

    if (this.players.has(this.playerId)) {
      this.players.get(this.playerId)!.setDirection(this.playDirection)
    }
    
    // send to server
    const message = new Message();
    message.setType(Message.Type.COMMAND);
    const direction = new Command();

    let xv: number = 0;
    let yv: number = 0;
    switch (this.playDirection) {
      case Direction.Down:
          yv = 1;
          xv = 0;
          break;
      case Direction.Up:
          yv = -1;
          xv = 0;
          break;
      case Direction.Left:
          yv = 0;
          xv = -1;
          break;
      case Direction.Right:
          yv = 0;
          xv = 1;
        break;
    }

    if (yv != 0 || xv != 0) {
      this.players.get(this.playerId)!.setVelocity(yv, xv)
    }
    
    direction.setTime(new Date().getTime())
    direction.setXv(xv);
    direction.setYv(yv);

    message.setData(direction.serializeBinary());
    this.communication.sendMessage(message)
    // end

    // if (this.players.has(this.playerId)) {
    //   const player = this.players.get(this.playerId)!
    //   player.addCommand(direction);
    // }

    // this.previousPlayDirection = this.playDirection;
    this.playDirection = Direction.None;
  }

  private updateEntities(dt: number) {
    // Update the player sprite animation
    this.players.forEach((player: Player) => {
      // player.interpolate(this.serverDelta)
      player.update(dt, 5000, 3000)

      // Collision check
      this.players.forEach((otherPlayer: Player) => {
        if(player.id != otherPlayer.id 
            && player.isNear(otherPlayer, this.playerNearValue) 
            && player.overlaps(otherPlayer)) {
              player.revertDirection();
              otherPlayer.revertDirection();
        }
      })
    });

    this.camera.update();
  }

  private render() { 
    this.canvasWidth = this.ctx.canvas.width;
    this.canvasHeight = this.ctx.canvas.height;

    this.ctx.clearRect(0, 0, this.canvasWidth, this.canvasHeight);

    this.map.draw(this.ctx, this.camera.xView, this.camera.yView);

    this.infoEl.render(this.ctx, this.canvasWidth, this.canvasHeight);

    this.players.forEach((player) => {
      player.render(this.ctx, this.camera.xView, this.camera.yView);
    });
  }

  private requestAnimationFrame = (callback: FrameRequestCallback) => {
    return window.requestAnimationFrame(callback);
  };

  private resizeCanvasToDisplaySize = (canvas: HTMLCanvasElement) => {
    // look up the size the canvas is being displayed
    const width = canvas.clientWidth;
    const height = canvas.clientHeight;

    // If it's resolution does not match change it
    if (canvas.width !== width || canvas.height !== height) {
      canvas.width = width;
      canvas.height = height;
    }
  };
}
