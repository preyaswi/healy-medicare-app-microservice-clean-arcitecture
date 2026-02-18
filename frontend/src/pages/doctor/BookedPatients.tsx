import { useEffect, useState } from 'react';
import api from '../../api/axios';
import { Patient } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import { User } from 'lucide-react';

export default function BookedPatients() {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    api.get('/doctor/patient')
      .then((res) => setPatients(res.data.data || []))
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <LoadingSpinner />;

  return (
    <div>
      <h1 className="page-title text-3xl mb-6">patient's list:</h1>

      {patients.length === 0 ? (
        <div className="card-yellow text-center py-12">
          <p className="text-gray-500 font-handwritten text-lg">No booked patients yet</p>
        </div>
      ) : (
        <div className="space-y-4">
          {patients.map((p) => (
            <div key={p.BookingId} className="card-yellow flex items-center gap-4">
              <div className="w-14 h-14 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center flex-shrink-0">
                <User className="h-7 w-7 text-gray-400" />
              </div>
              <div className="flex-1 min-w-0">
                <p className="font-handwritten font-bold text-base">{p.Fullname || 'patient-name'}</p>
                <p className="text-xs text-gray-600 font-handwritten">patient details:</p>
                <p className="text-xs text-gray-500 font-sans">{p.Email}</p>
                <p className="text-xs text-gray-500 font-sans">{p.Gender}</p>
              </div>
              <span className={`btn-dark text-xs py-2 px-5 flex-shrink-0 ${
                p.PaymentStatus === 'paid' ? 'bg-green-800' : ''
              }`}>
                {p.PaymentStatus === 'paid' ? 'paid' : 'pending'}
              </span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
