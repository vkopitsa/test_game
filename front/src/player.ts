import { Direction } from './game';

export class Player {

    private vx = 0;
    private vy = 0;
    private direction: Direction = Direction.Stop;

  constructor(
    public x = 100,
    public y = 100,
    public readonly speed = 75,
    private readonly radius = 5,
    private readonly color = 'blue'
  ) {
    this.vx = x;
    this.vy = y;
  }

  public update(dt: number, canvasWidth: number, canvasHeight: number) { 
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

  public setDirection(direction: Direction){
    this.direction = direction;
  }

  public render(ctx: CanvasRenderingContext2D) {
    ctx.beginPath();

    // Draws a ball
    ctx.arc(
      this.x,
      this.y,
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