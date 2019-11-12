
export class Map {
    // map texture
    image: HTMLImageElement | null = null;

    constructor(
        // map dimensions
        private width: number, private height: number
    ) {}

    // creates a prodedural generated map
    public generate() {
        var ctx = document.createElement("canvas").getContext("2d");
        if (!ctx) {
            return 
        }
        ctx.canvas.width = this.width;
        ctx.canvas.height = this.height;

        this.drawGrid(ctx)

        // store the generate map as this image texture
        this.image = new Image();
        this.image.src = ctx.canvas.toDataURL("image/png");

        // clear context
        ctx = null;
    }

    private drawGrid(
        ctx: CanvasRenderingContext2D, 
        minor: number | undefined = undefined, 
        major: number | undefined = undefined, 
        stroke: string | undefined = undefined, 
        fill: string | undefined = undefined,
    ){
        minor = minor || 10;
        major = major || minor * 5;
        stroke = stroke || "#00FF00";
        fill = fill || "#009900";
        ctx.save();
        ctx.strokeStyle = stroke;
        ctx.fillStyle = fill;
        let width = ctx.canvas.width, 
            height = ctx.canvas.height;
        
        for(var x = 0; x < width; x += minor) {
            ctx.beginPath();
            ctx.moveTo(x, 0);
            ctx.lineTo(x, height);
            ctx.lineWidth = (x % major == 0) ? 0.5 : 0.25;
            ctx.stroke();
            // coordinate
            // if(x % major == 0 ) {
            //     ctx.fillText(x.toString(), x, 10);
            // }
        }
  
        for(var y = 0; y < height; y += minor) {
            ctx.beginPath();
            ctx.moveTo(0, y);
  
            ctx.lineTo(width, y);
            ctx.lineWidth = (y % major == 0) ? 0.5 : 0.25;
            ctx.stroke();
            // coordinate
            // if(y % major == 0 ) {
            //     ctx.fillText(y.toString(), 0, y + 10);
            // }
        }
        ctx.restore();
    }

    // draw the map adjusted to camera
    public draw(context: CanvasRenderingContext2D, xView: number, yView: number) {
        if (!this.image) {
            return;
        }

        var sx, sy, dx, dy;
        var sWidth, sHeight, dWidth, dHeight;

        // offset point to crop the image
        sx = xView;
        sy = yView;

        // dimensions of cropped image			
        sWidth = context.canvas.width;
        sHeight = context.canvas.height;

        // if cropped image is smaller than canvas we need to change the source dimensions
        if (this.image.width - sx < sWidth) {
            sWidth = this.image.width - sx;
        }
        if (this.image.height - sy < sHeight) {
            sHeight = this.image.height - sy;
        }

        // location on canvas to draw the croped image
        dx = 0;
        dy = 0;
        // match destination with source to not scale the image
        dWidth = sWidth;
        dHeight = sHeight;

        context.drawImage(this.image, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight);

        // add current coordinate 
        const minor = 10;
        const major = minor * 5;
        const stroke = "#00FF00";
        const fill = "#009900";
        context.save();
        context.strokeStyle = stroke;
        context.fillStyle = fill;
        let width = context.canvas.width, 
            height = context.canvas.height;
        
        for(var x = 0; x < width; x += minor) {
            if(x % major == 0 ) {
                context.fillText((~~(xView + x)).toString(), x, 10);
            }
        }
  
        for(var y = 0; y < height; y += minor) {
            if(y % major == 0 ) {
                context.fillText((~~(yView + y)).toString(), 0, y + 10);
            }
        }
        context.restore();
    }
}