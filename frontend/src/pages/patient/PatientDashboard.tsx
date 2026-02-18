import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import { User } from 'lucide-react';
import api from '../../api/axios';
import { DoctorsDetails } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';

export default function PatientDashboard() {
  const { user } = useAuth();
  const [doctors, setDoctors] = useState<DoctorsDetails[]>([]);
  const [booked, setBooked] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Promise.all([
      api.get('/patient/doctor').then((res) => setDoctors(res.data.data || [])).catch(() => {}),
      api.get('/patient/booking').then((res) => setBooked(res.data.data || [])).catch(() => {}),
    ]).finally(() => setLoading(false));
  }, []);

  if (loading) return <LoadingSpinner />;

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="page-header">
        <h1 className="page-title text-4xl">welcome, {user?.name || 'patient'}</h1>
        <span className="brand-name text-2xl">LifeLink</span>
      </div>

      {/* Next Schedule - gray card */}
      <div className="card-gray py-8 px-6">
        <h2 className="font-handwritten text-xl font-bold mb-6">next schedule:</h2>
        {booked.length === 0 ? (
          <p className="text-gray-500 font-handwritten text-base">No upcoming schedules</p>
        ) : (
          <div className="flex gap-6 overflow-x-auto pb-2">
            {booked.slice(0, 4).map((b: any, i: number) => (
              <div key={i} className="flex flex-col items-center min-w-[120px]">
                <div className="w-16 h-16 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center mb-2">
                  <User className="h-8 w-8 text-gray-400" />
                </div>
                <p className="font-handwritten text-base italic">{b.DoctorName || 'doctor-name'}</p>
                <Link to={`/patient/doctor/${b.DoctorId || ''}`} className="btn-dark text-xs py-1.5 px-4 mt-2">
                  click here
                </Link>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Doctors List - yellow cards */}
      <div className="space-y-4">
        {doctors.map((doc) => (
          <div key={doc.DoctorDetail.Id} className="card-yellow flex items-center gap-4">
            <div className="w-16 h-16 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center flex-shrink-0">
              <User className="h-8 w-8 text-gray-400" />
            </div>
            <div className="flex-1 min-w-0">
              <p className="font-handwritten italic text-base">{doc.DoctorDetail.FullName || 'doctor-name'}</p>
              <p className="font-handwritten text-base font-bold mt-1">doctor's details:</p>
              <p className="text-xs text-gray-600 font-sans">Specialization: {doc.DoctorDetail.Specialization}</p>
              <p className="text-xs text-gray-600 font-sans">Years of Experience: {doc.DoctorDetail.YearsOfExperience}</p>
              <div className="flex gap-0.5 my-1">
                {[1, 2, 3, 4, 5].map((s) => (
                  <span key={s} className={`text-sm ${s <= (doc.Rating || 0) ? 'text-yellow-400' : 'text-gray-300'}`}>&#9733;</span>
                ))}
              </div>
              <p className="font-handwritten text-base font-bold">{doc.DoctorDetail.Fees ? 'active' : 'leave'}</p>
            </div>
            <Link to={`/patient/doctor/${doc.DoctorDetail.Id}`} className="btn-dark text-sm py-2.5 px-6 flex-shrink-0">
              Get to know
            </Link>
          </div>
        ))}
        {doctors.length === 0 && (
          <div className="card-yellow text-center py-12">
            <p className="text-gray-500 font-handwritten text-lg">No doctors available</p>
          </div>
        )}
      </div>

      {/* Pagination */}
      {doctors.length > 0 && (
        <div className="flex justify-center gap-2 font-handwritten text-lg italic text-gray-500">
          <span>1</span> <span>2</span> <span>3</span> <span>4</span> <span>&gt;</span>
        </div>
      )}
    </div>
  );
}
