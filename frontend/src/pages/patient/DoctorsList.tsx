import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import api from '../../api/axios';
import { DoctorsDetails } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import { User } from 'lucide-react';

export default function DoctorsList() {
  const [doctors, setDoctors] = useState<DoctorsDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');

  useEffect(() => {
    api.get('/patient/doctor')
      .then((res) => setDoctors(res.data.data || []))
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  const filtered = doctors.filter((d) =>
    d.DoctorDetail.FullName.toLowerCase().includes(search.toLowerCase()) ||
    d.DoctorDetail.Specialization.toLowerCase().includes(search.toLowerCase())
  );

  if (loading) return <LoadingSpinner />;

  return (
    <div className="space-y-6">
      <div className="page-header">
        <h1 className="page-title text-3xl">find a doctor</h1>
        <span className="brand-name text-xl">LifeLink</span>
      </div>

      {/* Search */}
      <div>
        <input type="text" className="input-field max-w-sm border border-gray-200"
          placeholder="Search by name or specialization..."
          value={search} onChange={(e) => setSearch(e.target.value)} />
      </div>

      {/* Doctors List - yellow cards */}
      {filtered.length === 0 ? (
        <div className="card-yellow text-center py-12">
          <p className="text-gray-500 font-handwritten text-lg">No doctors found</p>
        </div>
      ) : (
        <div className="space-y-4">
          {filtered.map((doc) => (
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
        </div>
      )}
    </div>
  );
}
