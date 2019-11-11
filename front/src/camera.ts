import { Rectangle } from './rectangle';
import { Player } from './player';

export enum Axis {
    NONE = 1,
    HORIZONTAL,
    VERTICAL,
    BOTH,
  }

export class Camera {
    // distance from followed object to border before camera starts move
    xDeadZone: number = 0; // min distance to horizontal borders
    yDeadZone: number = 0; // min distance to vertical borders
    // viewport dimensions
    wView: number;
    hView: number;
    // allow camera to move in vertical and horizontal axis
    axis: Axis = Axis.BOTH;
    // object that should be followed
    followed: Player | null = null;
    // rectangle that represents the viewport
    viewportRect: Rectangle;
    // rectangle that represents the world's boundary (room's boundary)
    worldRect: Rectangle;

    constructor(
        // position of camera (left-top coordinate)
        public xView: number = 0, public yView: number = 0, 
        viewportWidth: number, viewportHeight: number, 
        worldWidth: number, worldHeight: number
    ) {
        // viewport dimensions
        this.wView = viewportWidth;
        this.hView = viewportHeight;

        // rectangle that represents the viewport
        this.viewportRect = new Rectangle(this.xView, this.yView, this.wView, this.hView);

        // rectangle that represents the world's boundary (room's boundary)
        this.worldRect = new Rectangle(0, 0, worldWidth, worldHeight);
    }

    public follow(gameObject: any, xDeadZone: number, yDeadZone: number) {
        this.followed = gameObject;
        this.xDeadZone = xDeadZone;
        this.yDeadZone = yDeadZone;
    }

    public update() {
        // keep following the player (or other desired object)
        if (this.followed != null) {
            if (this.axis == Axis.HORIZONTAL || this.axis == Axis.BOTH) {
                // moves camera on horizontal axis based on followed object position
                if (this.followed.x - this.xView + this.xDeadZone > this.wView)
                    this.xView = this.followed.x - (this.wView - this.xDeadZone);
                else if (this.followed.x - this.xDeadZone < this.xView)
                    this.xView = this.followed.x - this.xDeadZone;

            }
            if (this.axis == Axis.VERTICAL || this.axis == Axis.BOTH) {
                // moves camera on vertical axis based on followed object position
                if (this.followed.y - this.yView + this.yDeadZone > this.hView)
                    this.yView = this.followed.y - (this.hView - this.yDeadZone);
                else if (this.followed.y - this.yDeadZone < this.yView)
                    this.yView = this.followed.y - this.yDeadZone;
            }

        }

        // update viewportRect
        this.viewportRect.set(this.xView, this.yView);

        // don't let camera leaves the world's boundary
        if (!this.viewportRect.within(this.worldRect)) {
            if (this.viewportRect.getLeft() < this.worldRect.getLeft())
                this.xView = this.worldRect.getLeft();
            if (this.viewportRect.getTop() < this.worldRect.getTop())
                this.yView = this.worldRect.getTop();
            if (this.viewportRect.getRight() > this.worldRect.getRight())
                this.xView = this.worldRect.getRight() - this.wView;
            if (this.viewportRect.getBottom() > this.worldRect.getBottom())
                this.yView = this.worldRect.getBottom() - this.hView;
        }

    }
    
}
