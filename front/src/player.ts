import { Direction } from './game';
import { Command } from './proto/command_pb';
import { Data } from './proto/data_pb';

export class Player {

    private vx = 0;
    private vy = 0;
    private direction: Direction = Direction.Stop;
    private commands: Command[] = [];
    private dataBuffer: [number, Data][] = [];

  constructor(
    public readonly id: any,
    public x = 100,
    public y = 100,
    public readonly speed = 200,
    private readonly radius = 50,
    private color = '#ffffff00'
  ) {
  }

  public update(dt: number, worldWidth: number, worldHeight: number){
    // parameter step is the time between frames ( in seconds )
    if (this.vy != 0) {
      this.y = this.y + (dt *this.speed * this.vy);
    }

    if (this.vx != 0) {
      this.x = this.x + (dt *this.speed * this.vx);
    }

    // don't let player leaves the world's boundary
    if (this.y + this.radius > worldHeight || this.y - this.radius < 0) {
      this.vy = (this.y + this.radius) > worldHeight ? -1 : 1;
    }

    if (this.x + this.radius > worldWidth || this.x - this.radius < 0) {
      this.vx = (this.x + this.radius) > worldWidth ? -1 : 1;
    }
}

public applyData(data: Data) {
  this.setPosition(data.getY(), data.getX())
  this.setVelocity(data.getYv(), data.getXv())
}

// reconciliation
public addCommand(command: Command) {
  this.commands.push(command);
}

public reconciliation(data: Data) {
  // Received the authoritative position of this player.
  // this.applyData(data);
  // this.y = data.getY()
  // this.x = data.getX()

  if (this.commands.length === 0) {
    this.applyData(data);
  } else if (data.getTime() !== 0) {
    var j = 0;
    while (j < this.commands.length) {
      var command = this.commands[j];
      if (command.getTime() <= data.getTime()) {
        // Already processed. Its effect is already taken into account into the world update
        // we just got, so we can drop it.
        this.commands.splice(j, 1);
      } else {
        // Not processed by the server yet. Re-apply it.
        this.applyData(data);
        j++;
      }
    }
  }
}

public interpolate(dt: number,) {
  // Compute render timestamp.
  var now = +new Date(); 
  var render_timestamp = now - (1000*dt);

  // console.log((dt))

  // if (this.dataBuffer.length >= 2){
  //   console.log(this.dataBuffer[1][0] <= render_timestamp)
  // }

  // Drop older positions.
  while (this.dataBuffer.length >= 2 && this.dataBuffer[1][0] <= render_timestamp) {
    this.dataBuffer.shift();
  }

  // console.log(this.dataBuffer.length);

  // Interpolate between the two surrounding authoritative positions.
  if (this.dataBuffer.length >= 2 && this.dataBuffer[0][0] <= render_timestamp && render_timestamp <= this.dataBuffer[1][0]) {
    var p0 = this.dataBuffer[0][1];
    var p1 = this.dataBuffer[1][1];
    var t0 = this.dataBuffer[0][0];
    var t1 = this.dataBuffer[1][0];

    // const x = p0.getX() + (p1.getX() - p0.getX()) * (render_timestamp - t0) / (t1 - t0);
    // const y = p0.getY() + (p1.getY() - p0.getY()) * (render_timestamp - t0) / (t1 - t0);
    // this.x = this.x + (dt *this.speed * this.vx);
    const x = p0.getX() + (p1.getX() - p0.getX()) * (render_timestamp - t0) / (t1 - t0)
    const y = p0.getY() + (p1.getY() - p0.getY()) * (render_timestamp - t0) / (t1 - t0)
    // this.applyData(this.dataBuffer[0][1])
    //console.log(x, y);
    this.setPosition(y, x)
    this.setVelocity(p0.getYv(), p0.getXv())
  }
}

public addData(data: Data) {
  var timestamp = +new Date();
  this.dataBuffer.push([timestamp, data]);
}

  public setColor(color: string){
    this.color = color;
  }

  public setDirection(direction: Direction){
    this.direction = direction;
  }

  public setPosition(y: number, x: number){
    this.y = y;
    this.x = x;
  }

  public setVelocity(vy: number, vx: number){
    this.vy = vy;
    this.vx = vx;
  }

  public render(ctx: CanvasRenderingContext2D, xView: number, yView: number) {
    ctx.beginPath();

    ctx.arc(
      this.x - xView,
      this.y - yView,
      this.radius,
      0,
      Math.PI * 2,
      true
    );

    ctx.closePath();

    // Colors and fills the ball
    ctx.fillStyle = this.color;
    ctx.fill();
  }

  // Determines is a player is near other player.
  public isNear(other: Player, playerNearValue: number): boolean { 
    const xdiff = Math.abs(this.x - other.x);
    const ydiff =Math.abs(this.y - other.y);
    const mdiff = Math.max(xdiff, ydiff);

    return mdiff < playerNearValue
  }

  public overlaps(other: Player): boolean { 
    const dx = this.x - other.x;
    const dy = this.y - other.y;
    const distance = Math.sqrt(dx*dx + dy*dy);

    return distance < this.radius+other.radius;
  }

  public revertDirection() { 
    this.setVelocity(-this.vy, -this.vx);
  }
}