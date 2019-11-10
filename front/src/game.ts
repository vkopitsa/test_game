import { IСommunication } from "./communication";
import { Message } from "./proto/message_pb";
import { Data } from "./proto/data_pb";
import { Direction as PDirection } from "./proto/direction_pb";
import { Player } from './player';

export enum Direction {
  Stop = 0,
  Up,
  Down,
  Left,
  Right,
}

export class Game {
  private readonly ctx: CanvasRenderingContext2D; // HTML Canvas's 2D context
  private readonly canvasWidth: number; // width of the canvas
  private readonly canvasHeight: number; // height of the canvas
  // private readonly ball = new Ball(50, 50); // create a new ball with x and y 50 and other properties default
  private lastTime: number = 0;
  private gameTime: number = 0;
  private playDirection: Direction = Direction.Stop;
  private previousPlayDirection: Direction = Direction.Stop;
  // private readonly player: Player = new Player(50, 50);
  private playerId: number = 0;
  private readonly players: Map<number,Player> = new Map<number,Player>();

  // Speed in pixels per second
  private playerSpeed = 100;

  constructor(canvas: HTMLCanvasElement, private communication: IСommunication) {
    this.ctx = canvas.getContext('2d')!;
    this.canvasWidth = canvas.width;
    this.canvasHeight = canvas.height;

    const that = this;
    window.addEventListener('keydown', function(event) {
      switch (event.keyCode) {
        case 37: // Left
          that.playDirection = Direction.Left;
        break;
    
        case 38: // Up
        that.playDirection = Direction.Up;
        break;
    
        case 39: // Right
        that.playDirection = Direction.Right;
        break;
    
        case 40: // Down
        that.playDirection = Direction.Down;
        break;
      }
    }, false);

    this.communication.onMesssage((msg: Message) => {
      if (msg.getType() === Message.Type.JOINED && !this.players.has(msg.getPlayerid())) {
        this.players.set(msg.getPlayerid(), new Player(50, 50))
        this.playerId = msg.getPlayerid();
      } else if (msg.getType() === Message.Type.DIRECTION) {
        const direction = PDirection.deserializeBinary(msg.getData_asU8());
        // this.players.forEach((player) => {
        //   player.setDirection(this.playDirection)
        // });
        if (this.players.has(msg.getPlayerid())) {
          this.players.get(msg.getPlayerid())!.setDirection(direction.getType())
        } else {
          const color = '#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6);
          const newPlayer = new Player(50, 50, 75, 5, color);
          newPlayer.setDirection(direction.getType())
          this.players.set(msg.getPlayerid(), newPlayer);
        }
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
    this.gameTime += dt;

    this.handleInput(dt);
    this.updateEntities(dt);

    // It gets harder over time by adding enemies using this
    // equation: 1-.993^gameTime
    // if(Math.random() < 1 - Math.pow(.993, gameTime)) {
    //     enemies.push({
    //         pos: [canvas.width,
    //               Math.random() * (canvas.height - 39)],
    //         sprite: new Sprite('img/sprites.png', [0, 78], [80, 39],
    //                            6, [0, 1, 2, 3, 2, 1])
    //     });
    // }

    // checkCollisions();

    // scoreEl.innerHTML = score;
  }

  private handleInput(dt: number) {

    // pass if the same direction
    if (this.previousPlayDirection === this.playDirection) {
      return
    }

    // this.players.forEach((player) => {
    //   player.setDirection(this.playDirection)
    // });
    if (this.players.has(this.playerId)) {
      this.players.get(this.playerId)!.setDirection(this.playDirection)
    }
    
    // send to server
    const message = new Message();
    message.setType(Message.Type.DIRECTION);
    const direction = new PDirection();
    direction.setType(this.playDirection)
    message.setData(direction.serializeBinary());
    this.communication.sendMessage(message)
    // end

    this.previousPlayDirection = this.playDirection;
  }

  private updateEntities(dt: number) {
    // Update the player sprite animation
    //this.player.update(dt, this.canvasWidth, this.canvasHeight);
    this.players.forEach((player) => {
      player.update(dt, this.canvasWidth, this.canvasHeight);
    });
  }

  private render() { 
    this.ctx.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
    // this.player.render(this.ctx);
    this.players.forEach((player) => {
      player.render(this.ctx);
    });
  }

  private requestAnimationFrame = (callback: FrameRequestCallback) => {
    return window.requestAnimationFrame(callback);
  };
}
