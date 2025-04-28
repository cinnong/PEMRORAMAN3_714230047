import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { LayoutDashboard } from "./components/layouts/LayoutDashboard";
import { Dashboard } from "./pages/Dashboard";
import { MahasiswaPage } from "./pages/MahasiswaPage";

function App() {
  return (
    <Router>
      <LayoutDashboard>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/mahasiswa" element={<MahasiswaPage />} />
        </Routes>
      </LayoutDashboard>
    </Router>
  );
}

export default App