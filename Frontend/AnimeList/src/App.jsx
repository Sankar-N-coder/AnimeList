import Signup from "./Componentrs/Signup"
import Home from "./Componentrs/Home"
import {BrowserRouter,Routes,Route} from "react-router-dom"
function App() {
  return(
    <>
    <BrowserRouter>
    <Routes>
        <Route path="/" element={<Signup />} />
        <Route path="/Animelist" element={<Home />} />
      </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
