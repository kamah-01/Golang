import './App.css'
import SampleForm from "./home/form";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <Router>
       <Routes>
          <Route path="/" element={<SampleForm/>} />
       </Routes>
    </Router>
  )
}

export default App;
