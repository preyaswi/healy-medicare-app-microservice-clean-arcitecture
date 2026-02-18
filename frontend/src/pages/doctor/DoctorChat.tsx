import { useEffect, useRef, useState } from 'react';
import { useAuth } from '../../context/AuthContext';
import { Send } from 'lucide-react';
import api from '../../api/axios';
import toast from 'react-hot-toast';

interface ChatMsg { sender_id: number; message_content: string; timestamp: string; }

export default function DoctorChat() {
  const { user } = useAuth();
  const [friendId, setFriendId] = useState('');
  const [connected, setConnected] = useState(false);
  const [messages, setMessages] = useState<ChatMsg[]>([]);
  const [newMessage, setNewMessage] = useState('');
  const wsRef = useRef<WebSocket | null>(null);
  const endRef = useRef<HTMLDivElement>(null);

  useEffect(() => { endRef.current?.scrollIntoView({ behavior: 'smooth' }); }, [messages]);

  const loadHistory = async () => {
    if (!friendId) return;
    try {
      const res = await api.get('/chat/messages', { params: { FriendID: friendId, Offset: '0', Limit: '50' } });
      if (res.data.data) setMessages(res.data.data);
    } catch {}
  };

  const connectChat = () => {
    if (!friendId) { toast.error('Enter a patient ID'); return; }
    loadHistory();
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    const ws = new WebSocket(`${protocol}//${window.location.host}/ws/chat?token=${user?.accessToken}`);
    ws.onopen = () => { setConnected(true); toast.success('Connected'); };
    ws.onmessage = (e) => {
      try { setMessages((prev) => [...prev, JSON.parse(e.data)]); }
      catch { setMessages((prev) => [...prev, { sender_id: parseInt(friendId), message_content: e.data, timestamp: new Date().toISOString() }]); }
    };
    ws.onclose = () => setConnected(false);
    ws.onerror = () => { toast.error('Connection error'); setConnected(false); };
    wsRef.current = ws;
  };

  const sendMessage = () => {
    if (!newMessage.trim() || !wsRef.current) return;
    wsRef.current.send(JSON.stringify({ RecipientID: friendId, Content: newMessage, TimeStamp: new Date().toISOString() }));
    setMessages((prev) => [...prev, { sender_id: parseInt(user?.id || '0'), message_content: newMessage, timestamp: new Date().toISOString() }]);
    setNewMessage('');
  };

  useEffect(() => { return () => { wsRef.current?.close(); }; }, []);

  return (
    <div className="max-w-2xl">
      <div className="flex items-center justify-between mb-6">
        <h1 className="font-handwritten text-3xl font-bold">chat</h1>
        <span className="brand-name text-xl">LifeLink</span>
      </div>

      {!connected ? (
        <div className="card-yellow py-8 px-6 text-center">
          <p className="font-handwritten text-lg mb-4">Start a conversation</p>
          <div className="flex gap-3 max-w-sm mx-auto">
            <input type="text" className="flex-1 px-4 py-2.5 bg-white rounded-xl border-none outline-none"
              placeholder="Patient ID..." value={friendId} onChange={(e) => setFriendId(e.target.value)} />
            <button onClick={connectChat} className="btn-dark text-sm">Connect</button>
          </div>
        </div>
      ) : (
        <div className="bg-white rounded-3xl border border-gray-100 flex flex-col h-[550px]">
          <div className="px-4 py-3 border-b border-gray-100 flex items-center justify-between">
            <div className="flex items-center gap-2">
              <div className="w-2 h-2 bg-green-500 rounded-full"></div>
              <span className="text-sm font-handwritten">Patient #{friendId}</span>
            </div>
            <button onClick={() => { wsRef.current?.close(); setConnected(false); }} className="text-xs text-red-500">Disconnect</button>
          </div>
          <div className="flex-1 overflow-y-auto p-4 space-y-3">
            {messages.length === 0 && <p className="text-center text-gray-400 text-sm py-8 font-handwritten">No messages yet</p>}
            {messages.map((msg, i) => {
              const isMe = msg.sender_id === parseInt(user?.id || '0');
              return (
                <div key={i} className={`flex ${isMe ? 'justify-end' : 'justify-start'}`}>
                  <div className={`max-w-[70%] px-4 py-2.5 rounded-2xl text-sm ${
                    isMe ? 'bg-brand-yellow rounded-br-md' : 'bg-brand-gray-light rounded-bl-md'
                  }`}>
                    {msg.message_content}
                    <div className="text-[10px] mt-1 text-gray-400">
                      {new Date(msg.timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                    </div>
                  </div>
                </div>
              );
            })}
            <div ref={endRef} />
          </div>
          <div className="px-4 py-3 border-t border-gray-100 flex gap-2">
            <input type="text" className="flex-1 px-4 py-2.5 bg-brand-gray-light rounded-xl outline-none text-sm"
              placeholder="Type a message..." value={newMessage} onChange={(e) => setNewMessage(e.target.value)}
              onKeyDown={(e) => e.key === 'Enter' && sendMessage()} />
            <button onClick={sendMessage} className="p-2.5 bg-brand-black text-white rounded-xl">
              <Send className="h-4 w-4" />
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
