import { useEffect, useState } from 'react';
import api from '../../api/axios';
import LoadingSpinner from '../../components/LoadingSpinner';
import toast from 'react-hot-toast';

interface PatientItem {
  Id: string;
  Fullname: string;
  Email: string;
  Gender: string;
  Contactnumber: string;
}

export default function PatientsList() {
  const [patients, setPatients] = useState<PatientItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [expandedId, setExpandedId] = useState<string | null>(null);

  useEffect(() => {
    api.get('/admin/dashboard/patients')
      .then((res) => setPatients(res.data.data || []))
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  const handleBlock = async (id: string) => {
    try {
      await api.patch(`/admin/patients/block/${id}`);
      toast.success('Patient blocked');
      setPatients(patients.filter((p) => p.Id !== id));
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Failed to block patient');
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
        <h2 className="font-handwritten text-xl font-bold tracking-wider">PATIENT'S LIST</h2>

        {patients.length === 0 ? (
          <p className="text-gray-500 font-handwritten text-base">No patients registered</p>
        ) : (
          <div className="space-y-4">
            {patients.map((p) => (
              <div key={p.Id} className="bg-white/60 rounded-2xl p-4">
                <div className="flex items-center justify-between">
                  <p className="font-handwritten font-bold text-base">{p.Fullname || 'N/A'}</p>
                  <button
                    onClick={() => setExpandedId(expandedId === p.Id ? null : p.Id)}
                    className="btn-dark text-xs py-1 px-4"
                  >
                    {expandedId === p.Id ? 'hide' : 'details'}
                  </button>
                </div>
                {expandedId === p.Id && (
                  <div className="mt-3 space-y-1">
                    <p className="font-handwritten text-base tracking-wider indent-4">PATIENT'S DETAILS</p>
                    <p className="text-sm text-gray-600 font-sans">Email: {p.Email}</p>
                    <p className="text-sm text-gray-600 font-sans">Gender: {p.Gender}</p>
                    <p className="text-sm text-gray-600 font-sans">Contact: {p.Contactnumber}</p>
                    <button onClick={() => handleBlock(p.Id)}
                      className="mt-2 font-handwritten text-base tracking-wider text-red-600 hover:underline">
                      DELETE PATIENT
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
