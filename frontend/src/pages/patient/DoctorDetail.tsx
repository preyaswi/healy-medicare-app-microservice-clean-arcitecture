import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from '../../api/axios';
import { IndDoctorDetail, GetAvailability } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import toast from 'react-hot-toast';
import { User, Star, CheckCircle } from 'lucide-react';
import { useAuth } from '../../context/AuthContext';

export default function DoctorDetail() {
  const { doctorId } = useParams();
  const { user } = useAuth();
  const [doctor, setDoctor] = useState<IndDoctorDetail | null>(null);
  const [slots, setSlots] = useState<GetAvailability[]>([]);
  const [loading, setLoading] = useState(true);
  const [date, setDate] = useState('');
  const [loadingSlots, setLoadingSlots] = useState(false);
  const [rating, setRating] = useState(0);
  const [submittingRating, setSubmittingRating] = useState(false);

  useEffect(() => {
    api.get(`/patient/doctor/${doctorId}`)
      .then((res) => setDoctor(res.data.data))
      .catch(() => toast.error('Failed to load doctor details'))
      .finally(() => setLoading(false));
  }, [doctorId]);

  const fetchSlots = async () => {
    if (!date) { toast.error('Please select a date'); return; }
    setLoadingSlots(true);
    try {
      const res = await api.get('/patient/doctor/availability', { params: { doctor_id: doctorId, date } });
      setSlots(res.data.data || []);
    } catch { toast.error('Failed to fetch availability'); }
    finally { setLoadingSlots(false); }
  };

  const bookSlot = (slotId: number) => {
    window.location.href = `/api/patient/bookdoctor?slot_id=${slotId}&patient_id=${user?.id}`;
  };

  const submitRating = async () => {
    if (rating < 1 || rating > 5) { toast.error('Select a rating 1-5'); return; }
    setSubmittingRating(true);
    try {
      await api.post(`/patient/doctor/rate/${doctorId}`, { rate: rating });
      toast.success('Rating submitted!');
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Failed to submit rating');
    } finally { setSubmittingRating(false); }
  };

  if (loading) return <LoadingSpinner />;
  if (!doctor) return <p className="text-gray-500 font-handwritten text-base">Doctor not found.</p>;

  return (
    <div className="space-y-6">
      <div className="flex items-end justify-end">
        <span className="brand-name text-2xl">LifeLink</span>
      </div>

      {/* Doctor Profile - yellow card */}
      <div className="card-yellow text-center py-10 px-8">
        <div className="w-24 h-24 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center mx-auto mb-3">
          <User className="h-12 w-12 text-gray-400" />
        </div>
        <p className="font-handwritten text-lg italic">{doctor.FullName || 'doctor-name'}</p>
        <div className="flex justify-center gap-0.5 my-2">
          {[1, 2, 3, 4, 5].map((i) => (
            <Star key={i} className="h-5 w-5 text-yellow-400" fill="currentColor" />
          ))}
        </div>
        <p className="font-handwritten text-base font-bold italic mt-2">doctor's details:</p>
        <div className="mt-4 space-y-2 text-sm font-sans text-gray-700">
          <p>Specialization: {doctor.Specialization}</p>
          <p>Years of Experience: {doctor.YearsOfExperience}</p>
          <p>for 1 hour: {doctor.Fees}</p>
        </div>
        <p className="font-handwritten text-2xl font-bold mt-6 italic">active</p>

        <button onClick={() => bookSlot(0)} className="btn-dark text-lg px-10 py-3 mt-4">
          pay online
        </button>
      </div>

      {/* Check Availability */}
      <div className="card-yellow py-8 px-6">
        <h2 className="font-handwritten text-xl font-bold mb-4">check availability</h2>
        <div className="flex flex-col sm:flex-row gap-3">
          <input type="date" className="input-field flex-1"
            value={date} onChange={(e) => setDate(e.target.value)} min={new Date().toISOString().split('T')[0]} />
          <button onClick={fetchSlots} disabled={loadingSlots} className="btn-dark text-sm px-6">
            {loadingSlots ? 'loading...' : 'check slots'}
          </button>
        </div>

        {slots.length > 0 && (
          <div className="mt-6 grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3">
            {slots.map((slot) => (
              <button key={slot.Slot_id} onClick={() => !slot.Is_booked && bookSlot(slot.Slot_id)}
                disabled={slot.Is_booked}
                className={`py-2.5 px-3 rounded-xl text-sm font-sans transition-colors ${
                  slot.Is_booked
                    ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
                    : 'bg-white text-gray-700 hover:bg-brand-yellow-light cursor-pointer'
                }`}>
                <div className="flex items-center justify-center gap-1.5">
                  {slot.Is_booked ? <span className="text-xs text-red-400">Booked</span> : <CheckCircle className="h-3.5 w-3.5 text-green-500" />}
                  <span>{slot.Time}</span>
                </div>
              </button>
            ))}
          </div>
        )}
      </div>

      {/* Rate Doctor */}
      <div className="card-yellow py-6 px-6">
        <h2 className="font-handwritten text-xl font-bold mb-4">give review:</h2>
        <div className="flex items-center gap-4">
          <div className="flex gap-1">
            {[1, 2, 3, 4, 5].map((val) => (
              <button key={val} onClick={() => setRating(val)} className="p-0.5 hover:scale-110 transition-transform">
                <Star className={`h-7 w-7 ${val <= rating ? 'text-yellow-400' : 'text-gray-300'}`}
                  fill={val <= rating ? 'currentColor' : 'none'} />
              </button>
            ))}
          </div>
          <button onClick={submitRating} disabled={submittingRating || rating === 0} className="btn-dark text-sm px-6">
            {submittingRating ? 'submitting...' : 'submit'}
          </button>
        </div>
      </div>
    </div>
  );
}
