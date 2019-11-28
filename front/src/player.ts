import { Direction } from './game';

export class Player {

    private vx = 0;
    private vy = 0;
    private direction: Direction = Direction.Stop;

  constructor(
    public readonly id: any,
    public x = 100,
    public y = 100,
    public readonly speed = 200,
    private readonly radius = 50,
    private color = '#ffffff00'
  ) {
    // this.vx = x;
    // this.vy = y;
  }

  public update222(dt: number, canvasWidth: number, canvasHeight: number) { 
    // Increments the ball's position using its velocity
    this.x = this.vx;
    this.y = this.vy;
    // this.x = this.speed*dt

    if(this.direction == Direction.Down) {
      this.vy += this.speed * dt;
    }

    if(this.direction == Direction.Up) {
      this.vy -= this.speed * dt;
    }

    if(this.direction == Direction.Left) {
      this.vx -= this.speed * dt;
    }

    if(this.direction == Direction.Right) {
      this.vx += this.speed * dt;
    }

    if (this.vy > canvasHeight || this.vy < 0) {
      this.direction = this.vy > canvasHeight ? Direction.Up : Direction.Down;
    }

    if (this.vx > canvasWidth || this.vx < 0) {
      this.direction = this.vx > canvasWidth ? Direction.Left : Direction.Right;
    }
  }

  public update2(dt: number, worldWidth: number, worldHeight: number){
    // parameter step is the time between frames ( in seconds )
    // console.log(this.direction);
    // check controls and move the player accordingly
    // if (this.direction == Direction.Left)
    //   this.x -= this.speed * dt;
    // if (this.direction == Direction.Up)
    //   this.y -= this.speed * dt;
    // if (this.direction == Direction.Right)
    //   this.x += this.speed * dt;
    // if (this.direction == Direction.Down)
    //   this.y += this.speed * dt;
    if (this.vy != 0) {
      this.y = this.y + (dt *this.speed * this.vy);
    }

    if (this.vx != 0) {
      this.x = this.x + (dt *this.speed * this.vx);
    }

    if (this.direction == Direction.None) {
      // if p.command.GetYv() != 0 {
      //   p.position.y = p.position.y + (dt * p.speed * p.command.GetYv())
      // }
      // if p.command.GetXv() != 0 {
      //   p.position.x = p.position.x + (dt * p.speed * p.command.GetXv())
      // }

      // if (this.vy != 0) {
      //   this.y = this.y + (dt *this.speed * this.vy);
      // }

      // if (this.vx != 0) {
      //   this.x = this.x + (dt *this.speed * this.vx);
      // }
    }

    // don't let player leaves the world's boundary
    // if (this.x - (this.radius * 2) / 2 < 0) {
    //   this.x = (this.radius * 2) / 2;
    // }
    // if (this.y - (this.radius * 2) / 2 < 0) {
    //   this.y = (this.radius * 2) / 2;
    // }
    // if (this.x + (this.radius * 2) / 2 > worldWidth) {
    //   this.x = worldWidth - (this.radius * 2) / 2;
    // }
    // if (this.y + (this.radius * 2) / 2 > worldHeight) {
    //   this.y = worldHeight - (this.radius * 2) / 2;
    // }

    if (this.y + this.radius > worldHeight || this.y - this.radius < -this.radius) {
      //this.direction = (this.y + this.radius) > worldHeight ? Direction.Up : Direction.Down;
      this.vy = (this.y + this.radius) > worldHeight ? -1 : 1;
    }

    if (this.x + this.radius > worldWidth || this.x - this.radius < -this.radius) {
      //this.direction = (this.x + this.radius) > worldWidth ? Direction.Left : Direction.Right;
      this.vx = (this.x + this.radius) > worldWidth ? -1 : 1;
    }

    // console.log(this.y, this.x);
    // console.log(this.vy, this.vx);
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

    // Draws a ball
    // ctx.arc(
    //   this.x,
    //   this.y,
    //   this.radius,
    //   0,
    //   Math.PI * 2,
    //   true
    // );
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
}