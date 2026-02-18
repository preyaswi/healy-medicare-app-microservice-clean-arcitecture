import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import api from '../../api/axios';
import { Patient } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import { User } from 'lucide-react';

export default function DoctorDashboard() {
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
    <div className="space-y-8">
      {/* Next Schedule - gray card */}
      <div className="card-gray">
        <h2 className="font-handwritten text-xl font-bold mb-4">next schedule:</h2>
        {patients.length > 0 ? (
          <div className="flex gap-6 overflow-x-auto pb-2">
            {patients.slice(0, 4).map((p) => (
              <div key={p.BookingId} className="flex flex-col items-center gap-2 min-w-[120px]">
                <div className="w-16 h-16 rounded-full bg-white border-2 border-gray-300 flex items-center justify-center">
                  <User className="h-8 w-8 text-gray-400" />
                </div>
                <span className="font-handwritten text-base text-center">{p.Fullname || 'patient-name'}</span>
                <Link to="/doctor/patients" className="btn-dark text-xs py-1.5 px-4">click here</Link>
              </div>
            ))}
          </div>
        ) : (
          <p className="text-gray-500 font-handwritten text-base">No upcoming schedules</p>
        )}
      </div>

      {/* Patient's list */}
      <div>
        <h2 className="font-handwritten text-2xl font-bold mb-4">patient's list:</h2>
        <div className="space-y-4">
          {patients.length === 0 ? (
            <div className="card-yellow text-center py-8">
              <p className="text-gray-500 font-handwritten text-lg">No booked patients yet</p>
            </div>
          ) : (
            patients.map((p) => (
              <div key={p.BookingId} className="card-yellow flex items-center gap-4">
                <div className="w-14 h-14 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center flex-shrink-0">
                  <User className="h-7 w-7 text-gray-400" />
                </div>
                <div className="flex-1 min-w-0">
                  <p className="font-handwritten text-base">{p.Fullname || 'patient-name'}</p>
                  <p className="text-xs text-gray-600 font-handwritten">patient details:</p>
                  <p className="text-xs text-gray-500 font-sans">{p.Email}</p>
                  <p className="text-xs text-gray-500 font-sans">{p.PaymentStatus || 'payment'}</p>
                </div>
                <span className={`btn-dark text-xs py-2 px-5 flex-shrink-0 ${
                  p.PaymentStatus === 'paid' ? 'bg-green-800' : ''
                }`}>
                  {p.PaymentStatus === 'paid' ? 'paid' : 'pending'}
                </span>
              </div>
            ))
          )}
        </div>
      </div>

      {/* Pagination */}
      {patients.length > 0 && (
        <div className="flex justify-center gap-3 font-handwritten text-lg italic text-gray-500">
          <span>1</span> <span>2</span> <span>3</span> <span>4</span> <span>&gt;</span>
        </div>
      )}
    </div>
  );
}
