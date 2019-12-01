
import { Point, Direction } from './game';

export class Control {
    constructor(
        private width: number,
        private height: number,
    ) {}

    /**
     * getDirectionByKeywordCode
     */
    public getDirectionBykeyCode(keyCode: number): Direction {
        let playDirection = Direction.None;
        switch (keyCode) {
            case 37: // Left
                playDirection = Direction.Left;
            break;
        
            case 38: // Up
                playDirection = Direction.Up;
            break;
        
            case 39: // Right
                playDirection = Direction.Right;
            break;
        
            case 40: // Down
                playDirection = Direction.Down;
            break;
            }
        return playDirection;
    }

    /**
     * -------
     * |\   /|
     * | \ / |
     * | / \ |
     * |/   \|
     * -------
     */
    public getDirectionByPoint(p: Point): Direction {
        // center
        const cx = this.width / 2;
        const cy = this.height / 2;

        const center = new Point(cx, cy)

        // - x
        // | y
        // [x, y]
        // top
        const topLeft = new Point(
            0,
            0,
        ); 
        const topRight = new Point(
            this.width,
            0,
        );

        // botton
        const bottonLeft = new Point(
            0,
            this.height,
        ); 
        const bottonRight = new Point(
            this.width,
            this.height,
        );

        // left
        const leftTop = new Point(
            this.width,
            0,
        ); 
        const leftBotton = new Point(
            this.width,
            this.height,
        );

        // right
        const rightTop = new Point(
            0,
            0,
        ); 
        const rightBotton = new Point(
            0,
            this.height,
        );

        // Up 
        if (this.pointInTriangle(p, topLeft, topRight, center)) {
            return Direction.Up;
        // Down
        } else if (this.pointInTriangle(p, center, bottonRight, bottonLeft)) {
            return Direction.Down;
        // Right
        } else if (this.pointInTriangle(p, leftTop, center, leftBotton)) {
            return Direction.Right;
        // Left
        } else if (this.pointInTriangle(p, rightTop, rightBotton, center)) {
            return Direction.Left;
        }
        return Direction.None;
    }

    private pointInTriangle(pt: Point, v1: Point, v2: Point, v3: Point): boolean {
        var d1: number, d2: number, d3: number;
        var has_neg: boolean, has_pos: boolean;

        d1 = this.sign(pt, v1, v2);
        d2 = this.sign(pt, v2, v3);
        d3 = this.sign(pt, v3, v1);

        has_neg = (d1 < 0) || (d2 < 0) || (d3 < 0);
        has_pos = (d1 > 0) || (d2 > 0) || (d3 > 0);

        return !(has_neg && has_pos);
    }

    private sign(p1: Point, p2: Point, p3: Point): number {
        return (p1.x - p3.x) * (p2.y - p3.y) - (p2.x - p3.x) * (p1.y - p3.y);
    }
}