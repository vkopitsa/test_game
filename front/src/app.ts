import './style/normalize.css';
import './style/style.css';
import { Message } from "./proto/message_pb";
import { Game } from './game';
import { Сommunication, IСommunication } from './communication';

class App {
    private communication: IСommunication = new Сommunication();

    constructor(private canvas: HTMLCanvasElement) {}

    public start() {
        new Game(this.canvas, this.communication).start();

        // join to game
        setTimeout(() => {
            const message = new Message();
            message.setType(Message.Type.JOIN);
            this.communication.sendMessage(message)
        }, 1000);
    }
}

const gameCanvas = <HTMLCanvasElement>document.getElementById('game');
new App(gameCanvas).start();