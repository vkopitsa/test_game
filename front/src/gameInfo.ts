import { Info } from "./proto/info_pb";

export class GameInfo {

  private info?: Info;

  constructor(
    public x = 100,
    public y = 100,
    private color = '#ffffff00'
  ) {
  }

  public setInfo(info: Info) { 
    this.info = info;
  }

  public render(ctx: CanvasRenderingContext2D, w: number, h: number) {

    if (!this.info){
      return;
    }

    const count = this.info.getCount();
    const players = this.info.getPlayersList();

    ctx.fillStyle = 'rgba(251, 251, 251, 0.3)';
    ctx.fillRect(w-85, 15, 80, 100);

    // display scores
    this.text(ctx, 'Players: ' + count, w-80, 30, 14, 'rgba(251, 251, 251, 0.76)');

    for(let i = 0; i < players.length; ++i){
      const player = players[i];
      const offset = i == 0 ? 15 : (i+1) * 15;
      this.text(
        ctx,
        '#' + (i+1) + ": " + player.getId() + " - " + player.getScore(),
        w-80,
        30+offset,
        14,
        'rgba(251, 251, 251, 0.76)',
      );
    }
  }

  private text(ctx: CanvasRenderingContext2D, text: string, x: number, y: number, size: number, col: string) {
      ctx.font = 'bold '+size+'px';
      ctx.fillStyle = col;
      ctx.fillText(text, x, y);
  }
}