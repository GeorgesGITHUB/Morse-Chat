import { Route, Routes, Navigate } from "react-router-dom"
import Login from "./Login"
import App from "./App"

function Router(){
    return(
        <Routes>
            <Route path="/" element={<Navigate to='/login'/>}/>
            <Route path="/login" element={<Login/>}/>
            <Route path="/app" element={<App/>}/>
            <Route path="*" element={<h1>404 - Page Not Found</h1>}/>
        </Routes>
    )
}

export default Router