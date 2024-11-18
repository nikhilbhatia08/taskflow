import "./App.css";
import SideBar from "./components/SideBar";
import Queues from "./components/Queues";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

function App() {
  return (
    <div className="flex">
      <Router>
        <SideBar />
        <Routes>
          <Route path="/Queues" element={<Queues />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
