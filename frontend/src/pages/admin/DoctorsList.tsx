import { useEffect, useState } from 'react';
import api from '../../api/axios';
import { DoctorsDetails } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import toast from 'react-hot-toast';

export default function AdminDoctorsList() {
  const [doctors, setDoctors] = useState<DoctorsDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [expandedId, setExpandedId] = useState<number | null>(null);

  useEffect(() => {
    api.get('/admin/dashboard/doctors')
      .then((res) => setDoctors(res.data.data || []))
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  const handleBlock = async (id: number) => {
    try {
      await api.patch(`/admin/doctors/block/${id}`);
      toast.success('Doctor blocked');
      setDoctors(doctors.filter((d) => d.DoctorDetail.Id !== id));
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Failed to block doctor');
    }
  };

  if (loading) return <LoadingSpinner />;

  return (
    <div className="space-y-6">
      <div className="page-header">
        <h1 className="page-title text-3xl tracking-widest">DASHBOARD</h1>
        <span className="brand-name text-3xl">LifeLink</span>
      </div>

      <div className="card-yellow py-8 px-8 space-y-6">
        <h2 className="font-handwritten text-xl font-bold tracking-wider">DOCTOR'S LIST</h2>

        {doctors.length === 0 ? (
          <p className="text-gray-500 font-handwritten text-base">No doctors registered</p>
        ) : (
          <div className="space-y-4">
            {doctors.map((doc) => (
              <div key={doc.DoctorDetail.Id} className="bg-white/60 rounded-2xl p-4">
                <div className="flex items-center justify-between">
                  <p className="font-handwritten font-bold text-base">Dr. {doc.DoctorDetail.FullName}</p>
                  <button
                    onClick={() => setExpandedId(expandedId === doc.DoctorDetail.Id ? null : doc.DoctorDetail.Id)}
                    className="btn-dark text-xs py-1 px-4"
                  >
                    {expandedId === doc.DoctorDetail.Id ? 'hide' : 'details'}
                  </button>
                </div>
                {expandedId === doc.DoctorDetail.Id && (
                  <div className="mt-3 space-y-1">
                    <p className="font-handwritten text-base tracking-wider indent-4">DOCTOR'S DETAILS</p>
                    <p className="text-sm text-gray-600 font-sans">Email: {doc.DoctorDetail.Email}</p>
                    <p className="text-sm text-gray-600 font-sans">Specialization: {doc.DoctorDetail.Specialization}</p>
                    <p className="text-sm text-gray-600 font-sans">Experience: {doc.DoctorDetail.YearsOfExperience} years</p>
                    <p className="text-sm text-gray-600 font-sans">Fee: {doc.DoctorDetail.Fees}</p>
                    <p className="font-handwritten text-base tracking-wider mt-2">TOTAL LEAVES: —</p>
                    <button onClick={() => handleBlock(doc.DoctorDetail.Id)}
                      className="mt-2 font-handwritten text-base tracking-wider text-red-600 hover:underline">
                      DELETE DOCTOR
                    </button>
                  </div>
                )}
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
