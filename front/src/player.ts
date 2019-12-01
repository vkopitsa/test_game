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
  }

  public update222(dt: number, canvasWidth: number, canvasHeight: number) { 
    // Increments the ball's position using its velocity
    this.x = this.vx;
    this.y = this.vy;

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
    if (this.vy != 0) {
      this.y = this.y + (dt *this.speed * this.vy);
    }

    if (this.vx != 0) {
      this.x = this.x + (dt *this.speed * this.vx);
    }

    // don't let player leaves the world's boundary
    if (this.y + this.radius > worldHeight || this.y - this.radius < -this.radius) {
      this.vy = (this.y + this.radius) > worldHeight ? -1 : 1;
    }

    if (this.x + this.radius > worldWidth || this.x - this.radius < -this.radius) {
      this.vx = (this.x + this.radius) > worldWidth ? -1 : 1;
    }
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
}