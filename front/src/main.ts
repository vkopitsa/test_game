import './style/normalize.css';
import './style/style.css';
import { Message } from "./proto/message_pb";

class Game {
    constructor() {
    }

    public start() {
        this.ws(this.url("ws"))
    }

    private url(s: string) {
        var l: Location = window.location;
        return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != "80") && (l.port != "443")) ? ":" + l.port : "") + l.pathname + s;
    }

    private ws(url: string) {
        var ws = new WebSocket(url);
                
        var message = new Message();

        ws.onopen = function() {
            ws.binaryType = "arraybuffer";
            // message.setType(Message.Type.HELLO);
            // message.setData("");
            // ws.send(message.serializeBinary());

            // console.log("Connection is opened...");
        };

        ws.onmessage = function (event) { 
            var message = proto.messages.Message.deserializeBinary(event.data);
            if (message.getType() === proto.messages.Message.Type.DATA) {

                var data = proto.messages.Data.deserializeBinary(message.getData());
            }
        };
        
        ws.onclose = function() { 
            console.log("Connection is closed...");
        };
    }


}

new Game().start();