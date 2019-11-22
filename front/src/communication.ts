import { Message } from "./proto/message_pb";

export interface IСommunication {
    sendMessage(message: Message): void;
    onMesssage(handler: (message: Message) => void): void;
}

export class Сommunication implements IСommunication { 

    private connect: WebSocket;
    private handlers: { (message: Message): void; }[] = [];
    private connected: boolean = false;

    constructor() {
        this.connect = new WebSocket(this.url("ws"));
        
        this.setup();

        setInterval(() => {
            if (this.connect.CLOSED == this.connect.readyState) {
                this.connect = new WebSocket(this.url("ws"));
                this.setup();
            }
        }, 1000)
    }

    public sendMessage(message: Message): void {
        if (this.connected) {
            this.connect.send(message.serializeBinary());
        }
    }; 
    
    public onMesssage(handler: (message: Message) => void) {
        this.handlers.push(handler);
    }

    private setup() {
        const that = this;

        this.connect.onopen = function() {
            this.binaryType = "arraybuffer";
            that.connected = true;
        };

        this.connect.onclose = function() {
            that.connected = false;
        };

        this.connect.onmessage = function (event) { 
            const message = Message.deserializeBinary(event.data);
            for(const handler of that.handlers) {
                handler(message)
            }
        };
    }

    private url(s: string) {
        var l: Location = window.location;
        return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != "80") && (l.port != "443")) ? ":" + l.port : "") + l.pathname + s;
    }
}