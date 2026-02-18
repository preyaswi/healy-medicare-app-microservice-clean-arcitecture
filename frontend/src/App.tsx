import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Toaster } from 'react-hot-toast';
import { AuthProvider } from './context/AuthContext';
import Layout from './components/Layout';
import ProtectedRoute from './components/ProtectedRoute';

// Public pages
import Landing from './pages/Landing';
import GoogleCallback from './pages/patient/GoogleCallback';

// Patient auth (public)
import PatientLogin from './pages/patient/PatientLogin';
import PatientSignup from './pages/patient/PatientSignup';

// Doctor auth (public)
import DoctorLogin from './pages/doctor/DoctorLogin';
import DoctorSignup from './pages/doctor/DoctorSignup';

// Admin auth (public)
import AdminLogin from './pages/admin/AdminLogin';
import AdminSignup from './pages/admin/AdminSignup';

// Patient pages (protected)
import PatientDashboard from './pages/patient/PatientDashboard';
import PatientProfile from './pages/patient/PatientProfile';
import DoctorsList from './pages/patient/DoctorsList';
import DoctorDetail from './pages/patient/DoctorDetail';
import PatientChat from './pages/patient/PatientChat';

// Doctor pages (protected)
import DoctorDashboard from './pages/doctor/DoctorDashboard';
import DoctorProfileEdit from './pages/doctor/DoctorProfileEdit';
import SetAvailability from './pages/doctor/SetAvailability';
import BookedPatients from './pages/doctor/BookedPatients';
import CreatePrescription from './pages/doctor/CreatePrescription';
import DoctorChat from './pages/doctor/DoctorChat';

// Admin pages (protected)
import AdminDashboard from './pages/admin/AdminDashboard';
import PatientsList from './pages/admin/PatientsList';
import AdminDoctorsList from './pages/admin/DoctorsList';

export default function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Toaster position="top-right" toastOptions={{
          duration: 3000,
          style: { fontSize: '14px' },
        }} />
        <Routes>
          {/* Public routes without layout (Landing) */}
          <Route path="/" element={<Landing />} />
          <Route path="/google/redirect" element={<GoogleCallback />} />

          {/* Public routes with layout */}
          <Route element={<Layout />}>
            <Route path="/patient/login" element={<PatientLogin />} />
            <Route path="/patient/signup" element={<PatientSignup />} />
            <Route path="/doctor/login" element={<DoctorLogin />} />
            <Route path="/doctor/signup" element={<DoctorSignup />} />
            <Route path="/admin/login" element={<AdminLogin />} />
            <Route path="/admin/signup" element={<AdminSignup />} />

            {/* Patient protected routes */}
            <Route path="/patient/dashboard" element={<ProtectedRoute role="patient"><PatientDashboard /></ProtectedRoute>} />
            <Route path="/patient/profile" element={<ProtectedRoute role="patient"><PatientProfile /></ProtectedRoute>} />
            <Route path="/patient/doctors" element={<ProtectedRoute role="patient"><DoctorsList /></ProtectedRoute>} />
            <Route path="/patient/doctor/:doctorId" element={<ProtectedRoute role="patient"><DoctorDetail /></ProtectedRoute>} />
            <Route path="/patient/chat" element={<ProtectedRoute role="patient"><PatientChat /></ProtectedRoute>} />

            {/* Doctor protected routes */}
            <Route path="/doctor/dashboard" element={<ProtectedRoute role="doctor"><DoctorDashboard /></ProtectedRoute>} />
            <Route path="/doctor/profile" element={<ProtectedRoute role="doctor"><DoctorProfileEdit /></ProtectedRoute>} />
            <Route path="/doctor/availability" element={<ProtectedRoute role="doctor"><SetAvailability /></ProtectedRoute>} />
            <Route path="/doctor/patients" element={<ProtectedRoute role="doctor"><BookedPatients /></ProtectedRoute>} />
            <Route path="/doctor/prescriptions" element={<ProtectedRoute role="doctor"><CreatePrescription /></ProtectedRoute>} />
            <Route path="/doctor/chat" element={<ProtectedRoute role="doctor"><DoctorChat /></ProtectedRoute>} />

            {/* Admin protected routes */}
            <Route path="/admin/dashboard" element={<ProtectedRoute role="admin"><AdminDashboard /></ProtectedRoute>} />
            <Route path="/admin/patients" element={<ProtectedRoute role="admin"><PatientsList /></ProtectedRoute>} />
            <Route path="/admin/doctors" element={<ProtectedRoute role="admin"><AdminDoctorsList /></ProtectedRoute>} />
          </Route>
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  );
}
