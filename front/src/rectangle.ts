
export class Rectangle {
    private right: number;
    private bottom: number;

    constructor(
        private left: number = 0, 
        private top: number = 0, 
        private width: number = 0, 
        private height: number = 0,
    ) {
        this.right = this.left + this.width;
        this.bottom = this.top + this.height;
    }

    public set(
        left: number, 
        top: number, 
        width: number | undefined = undefined, 
        height: number | undefined = undefined,
    ) {
        this.left = left;
        this.top = top;
        this.width = width || this.width;
        this.height = height || this.height
        this.right = (this.left + this.width);
        this.bottom = (this.top + this.height);
    }

    public within(r: Rectangle){
        return (r.left <= this.left &&
        r.right >= this.right &&
        r.top <= this.top &&
        r.bottom >= this.bottom);
    }

    public overlaps(r: Rectangle){
        return (this.left < r.right &&
        r.left < this.right &&
        this.top < r.bottom &&
        r.top < this.bottom);
    }
    
    public getTop(): number {
        return this.top;
    }

    public getBottom(): number {
        return this.bottom;
    }

    public getLeft(): number {
        return this.left;
    }

    public getRight(): number {
        return this.right;
    }


}