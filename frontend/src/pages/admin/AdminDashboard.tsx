import { useEffect, useState } from 'react';
import api from '../../api/axios';

export default function AdminDashboard() {
  const [patientCount, setPatientCount] = useState(0);
  const [doctorCount, setDoctorCount] = useState(0);
  const [activeTab, setActiveTab] = useState('weekly');

  useEffect(() => {
    api.get('/admin/dashboard/patients')
      .then((res) => {
        const data = res.data.data;
        setPatientCount(Array.isArray(data) ? data.length : 0);
      })
      .catch(() => {});
    api.get('/admin/dashboard/doctors')
      .then((res) => {
        const data = res.data.data;
        setDoctorCount(Array.isArray(data) ? data.length : 0);
      })
      .catch(() => {});
  }, []);

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="page-header">
        <h1 className="page-title text-3xl tracking-widest">DASHBOARD</h1>
        <span className="brand-name text-3xl">LifeLink</span>
      </div>

      {/* Main yellow card */}
      <div className="card-yellow py-10 px-8">
        {/* Tabs */}
        <div className="flex justify-end gap-6 mb-8">
          {['WEEKLY', 'MONTHLY', 'YEARLY'].map((tab) => (
            <button
              key={tab}
              onClick={() => setActiveTab(tab.toLowerCase())}
              className={`font-handwritten text-base tracking-wider ${
                activeTab === tab.toLowerCase() ? 'font-bold text-brand-black' : 'text-gray-500'
              }`}
            >
              {tab}
            </button>
          ))}
        </div>

        {/* Counts */}
        <div className="space-y-6">
          <div>
            <span className="font-handwritten text-xl font-bold tracking-wider">DOCTOR:</span>
            <span className="ml-4 text-xl font-sans">{doctorCount}</span>
          </div>
          <div>
            <span className="font-handwritten text-xl font-bold tracking-wider">PATIENT:</span>
            <span className="ml-4 text-xl font-sans">{patientCount}</span>
          </div>
        </div>
      </div>
    </div>
  );
}
