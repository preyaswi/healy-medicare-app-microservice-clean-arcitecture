// API Response wrapper
export interface ApiResponse<T = unknown> {
  message: string;
  data: T;
  error: string | null;
}

// Patient types
export interface GoogleSignupDetailResponse {
  Id: string;
  GoogleId: string;
  FullName: string;
  Email: string;
}

export interface TokenPatient {
  Patient: GoogleSignupDetailResponse;
  AccessToken: string;
  RefreshToken: string;
}

export interface PatientDetails {
  fullname: string;
  email: string;
  gender: string;
  contactnumber: string;
}

export interface Patient {
  BookingId: number;
  SlotId: number;
  PaymentStatus: string;
  PatientId: number;
  Fullname: string;
  Email: string;
  Gender: string;
  Contactnumber: string;
}

// Doctor types
export interface DoctorSignUp {
  full_name: string;
  email: string;
  phone_number: string;
  password: string;
  confirm_password: string;
  specialization: string;
  years_of_experience: number;
  license_number: string;
  fees: number;
}

export interface DoctorLogin {
  email: string;
  password: string;
}

export interface DoctorDetail {
  Id: number;
  FullName: string;
  Email: string;
  PhoneNumber: string;
  Specialization: string;
  YearsOfExperience: number;
  LicenseNumber: string;
  Fees: number;
}

export interface DoctorsDetails {
  DoctorDetail: DoctorDetail;
  Rating: number;
}

export interface IndDoctorDetail extends DoctorDetail {
  Rating: number;
}

export interface DoctorSignUpResponse {
  DoctorDetail: DoctorDetail;
  AccessToken: string;
  RefreshToken: string;
}

export interface DoctorDetailsUpdate {
  FullName: string;
  Email: string;
  PhoneNumber: string;
  Specialization: string;
  YearsOfExperience: number;
  Fees: number;
}

export interface Rate {
  rate: number;
}

// Admin types
export interface AdminLogin {
  email: string;
  password: string;
}

export interface AdminSignUp {
  firstname: string;
  lastname: string;
  email: string;
  password: string;
}

export interface AdminDetailsResponse {
  id: number;
  firstname: string;
  lastname: string;
  Email: string;
}

export interface TokenAdmin {
  Admin: AdminDetailsResponse;
  Token: string;
}

// Booking / Availability types
export interface SetAvailability {
  date: string;
  starttime: string;
  endtime: string;
}

export interface GetAvailability {
  Slot_id: number;
  Time: string;
  Is_booked: boolean;
}

export interface CombinedBookingDetails {
  BookingId: number;
  PatientName: string;
  DoctorId: number;
  DoctorName: string;
  DoctorEmail: string;
  Fees: number;
  PaymentStatus: string;
}

// Prescription types
export interface PrescriptionRequest {
  doctor_id: number;
  booking_id: number;
  medicine: string;
  dosage: string;
  notes: string;
}

export interface CreatedPrescription {
  id: number;
  doctor_id: number;
  patient_id: string;
  booking_id: number;
  medicine: string;
  dosage: string;
  notes: string;
}

// Chat types
export interface Message {
  SenderID: string;
  RecipientID: string;
  Content: string;
  TimeStamp: string;
}

export interface ChatMessage {
  id: string;
  sender_id: number;
  chat_id: string;
  seen: boolean;
  image: string;
  message_content: string;
  timestamp: string;
}

// Auth user type for context
export type UserRole = 'patient' | 'doctor' | 'admin';

export interface AuthUser {
  id: string;
  name: string;
  email: string;
  role: UserRole;
  accessToken: string;
  refreshToken?: string;
}
