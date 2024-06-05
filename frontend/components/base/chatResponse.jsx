"use client"

import { useEffect, useRef, useState } from "react"
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Textarea } from "../ui/textarea";
import { ScrollArea } from "../ui/scroll-area";

//chatResponse takes a prompt as a prop
export default function ChatResponse({ prompt }) {
    const [message, setMessage] = useState('');
    const [input, setInput] = useState(prompt);
    let socket = useRef(null);

    useEffect(() => {
        // Open a WebSocket connection to the Go server
        socket.current = new WebSocket('ws://localhost:8000/generate');

        socket.current.onopen = () => {
            console.log('Connected to server');
        };

        socket.current.onmessage = (event) => {
            const newMessage = event.data;
            setMessage((prevMessage) => prevMessage + newMessage);
            console.log('Received message:', event.data);
        };

        socket.current.onclose = () => {
            console.log('Disconnected from server');
        };

        socket.current.onerror = (error) => {
            console.error('WebSocket error:', error);
        };

        return () => {
            socket.current.close();
        };
    }, []);

    const sendMessage = () => {
        if (socket.current && socket.current.readyState === WebSocket.OPEN) {
            socket.current.send(input);
            setInput('');
            setMessage('');
        } else {
            console.log('WebSocket is not open');
        }
    };

    return (
        <div>
            <ScrollArea className="h-[200px] w-[350px] rounded-md border p-4">
                You'll see the response here:
                <div style={{ whiteSpace: 'pre-wrap' }}>{message}</div>
            </ScrollArea>
            <Textarea value={input} onChange={(e) => setInput(e.target.value)} rows="8" />
            <Button onClick={sendMessage}>Send</Button>
        </div>
    )
}

