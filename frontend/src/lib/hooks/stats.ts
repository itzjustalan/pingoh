export const dd = {};
// import { useEffect, useRef, useState } from "react";
// import { env } from "../../env";
// import { authStore } from "../stores/auth";

// initiateConnection() {
//   const url = `ws://${env.backend.baseUrl}/api/stats/task/${this.taskId}`;
//   const token = authStore.getState().user?.access_token;
//   return new WebSocket(`${url}?token=${token}`);
// }

// export class HttpResults {
//   private connection: WebSocket | null;
//   private messageCallbacks: ((message: string) => void)[];
//   private openCallbacks: (() => void)[];
//   private closeCallbacks: ((event: CloseEvent) => void)[];
//   private errorCallbacks: ((error: Error) => void)[];
//
//   // constructor(url: string, options: WebSocketOptions = {})
//   constructor() {
//     this.connection = null;
//     this.messageCallbacks = [];
//     this.openCallbacks = [];
//     this.closeCallbacks = [];
//     this.errorCallbacks = [];
//     const url = `ws://${env.backend.baseUrl}/api/stats/task/${1}`;
//     const token = authStore.getState().user?.access_token;
//     this.connect(`${url}?token=${token}`);
//   }
//
//   connect(url: string, options: WebSocketOptions = {}) {
//     if (this.connection) {
//       console.warn("WebSocket connection already open");
//       return;
//     }
//
//     this.connection = new WebSocket(url);
//
//     this.connection.onopen = () => {
//       console.log("WebSocket connection opened");
//       this.openCallbacks.forEach((callback) => callback());
//     };
//
//     this.connection.onmessage = (event) => {
//       console.log("Received message:", event.data);
//       this.messageCallbacks.forEach((callback) => callback(event.data));
//     };
//
//     this.connection.onclose = (event) => {
//       console.log("WebSocket connection closed:", event);
//       this.connection = null;
//       this.closeCallbacks.forEach((callback) => callback(event));
//     };
//
//     this.connection.onerror = (error) => {
//       console.error("WebSocket error:", error);
//       this.errorCallbacks.forEach((callback) => callback(error));
//     };
//   }
//
//   disconnect() {
//     if (this.connection) {
//       this.connection.close();
//       this.connection = null;
//     }
//   }
//
//   sendMessage(message: string) {
//     if (this.connection && this.connection.readyState === WebSocket.OPEN) {
//       this.connection.send(message);
//     } else {
//       console.error("WebSocket connection is not open or ready");
//       // Handle sending messages when not connected (e.g., queue messages)
//     }
//   }
//
//   onMessage(callback: (message: string) => void) {
//     this.messageCallbacks.push(callback);
//   }
//
//   onOpen(callback: () => void) {
//     this.openCallbacks.push(callback);
//   }
//
//   onClose(callback: (event: CloseEvent) => void) {
//     this.closeCallbacks.push(callback);
//   }
//
//   onError(callback: (error: Error) => void) {
//     this.errorCallbacks.push(callback);
//   }
// }

// // Example usage:
// const wsManager = new WebSocketManager('wss://example.com');
//
// wsManager.onOpen(() => {
//   console.log('WebSocket connection established');
//   wsManager.sendMessage('Hello from the client');
// });
//
// wsManager.onMessage((message) => {
//   console.log('Received message:', message);
// });
//
// wsManager.onError((error) => {
//   console.error('WebSocket error:', error);
// });
//
// wsManager.connect();

/*
export const useWebhookConnection = (onMessage: (message: string) => void) => {
  const [isConnected, setIsConnected] = useState(false);
  const [connection, setConnection] = useState<null | WebSocket>();
  const messageHandlerRef = useRef(onMessage);

  useEffect(() => {
    const url = `ws://${env.backend.baseUrl}/api/stats/task/${1}`;
    const token = authStore.getState().user?.access_token;
    const ws = new WebSocket(`${url}?token=${token}`);

    ws.onopen = () => {
      console.log("Connected to websocket");
      setIsConnected(true);
    };

    ws.onmessage = (event) => {
      messageHandlerRef.current(event.data);
    };

    ws.onclose = () => {
      console.log("websocket closed");
      setIsConnected(false);
      setConnection(null);
    };

    setConnection(ws);

    return () => {
      ws.close();
    };
  }, []);

  const sendMessage = (message: string) => {
    if (connection) {
      connection.send(message);
    }
  };

  return {
    isConnected,
    sendMessage,
  };
};
*/
