import { useState } from 'react';
import api from '../../api/axios';
import toast from 'react-hot-toast';

export default function SetAvailability() {
  const [form, setForm] = useState({ date: '', starttime: '', endtime: '' });
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!form.date || !form.starttime || !form.endtime) {
      toast.error('Please fill all fields');
      return;
    }
    setLoading(true);
    try {
      const res = await api.post('/doctor/profile/availability', form);
      setResult(res.data.data);
      toast.success('Availability set!');
      setForm({ date: '', starttime: '', endtime: '' });
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Failed to set availability');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <form onSubmit={handleSubmit} className="card-yellow py-10 px-8 text-center space-y-6">
        <div>
          <input type="date" className="input-field max-w-sm mx-auto block text-center"
            value={form.date}
            onChange={(e) => setForm({ ...form, date: e.target.value })}
            min={new Date().toISOString().split('T')[0]}
            required />
        </div>

        <div className="flex items-center justify-center gap-4 max-w-sm mx-auto">
          <div className="flex-1">
            <label className="form-label text-center text-sm">Start</label>
            <input type="time" className="input-field text-center"
              value={form.starttime}
              onChange={(e) => setForm({ ...form, starttime: e.target.value })}
              required />
          </div>
          <span className="text-gray-400 pt-4 font-handwritten">-</span>
          <div className="flex-1">
            <label className="form-label text-center text-sm">End</label>
            <input type="time" className="input-field text-center"
              value={form.endtime}
              onChange={(e) => setForm({ ...form, endtime: e.target.value })}
              required />
          </div>
        </div>

        <button type="submit" disabled={loading} className="btn-dark text-lg px-12 py-3">
          {loading ? 'saving...' : 'schedule'}
        </button>
      </form>

      {result && (
        <div className="card-yellow">
          <h3 className="font-handwritten font-bold text-lg mb-2">Slots Created</h3>
          <pre className="text-sm font-sans text-gray-600 bg-white/60 p-3 rounded-xl overflow-auto">
            {JSON.stringify(result, null, 2)}
          </pre>
        </div>
      )}
    </div>
  );
}
