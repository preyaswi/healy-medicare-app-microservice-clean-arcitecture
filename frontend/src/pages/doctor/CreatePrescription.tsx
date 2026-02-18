import { useState } from 'react';
import api from '../../api/axios';
import toast from 'react-hot-toast';

export default function CreatePrescription() {
  const [form, setForm] = useState({ booking_id: '', medicine: '', dosage: '', notes: '' });
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!form.booking_id || !form.medicine || !form.dosage) {
      toast.error('Please fill required fields');
      return;
    }
    setLoading(true);
    try {
      await api.post(`/doctor/patient/prescription?booking_id=${form.booking_id}`, {
        medicine: form.medicine, dosage: form.dosage, notes: form.notes,
      });
      toast.success('Prescription created!');
      setForm({ booking_id: '', medicine: '', dosage: '', notes: '' });
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Failed to create prescription');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-lg">
      <h1 className="page-title text-3xl mb-6">prescription</h1>

      <form onSubmit={handleSubmit} className="card-yellow py-8 px-6 sm:px-10 space-y-5">
        <div>
          <label className="form-label">Booking ID</label>
          <input type="number" className="input-field"
            value={form.booking_id} onChange={(e) => setForm({ ...form, booking_id: e.target.value })} required />
        </div>
        <div>
          <label className="form-label">Medicine</label>
          <input type="text" className="input-field"
            value={form.medicine} onChange={(e) => setForm({ ...form, medicine: e.target.value })} required placeholder="e.g., Amoxicillin 500mg" />
        </div>
        <div>
          <label className="form-label">Dosage</label>
          <input type="text" className="input-field"
            value={form.dosage} onChange={(e) => setForm({ ...form, dosage: e.target.value })} required placeholder="e.g., 1 tablet 3 times a day" />
        </div>
        <div>
          <label className="form-label">Notes</label>
          <textarea className="input-field min-h-[80px]"
            value={form.notes} onChange={(e) => setForm({ ...form, notes: e.target.value })} placeholder="Additional notes..." />
        </div>
        <button type="submit" disabled={loading} className="btn-dark px-8 py-3">
          {loading ? 'creating...' : 'submit'}
        </button>
      </form>
    </div>
  );
}
