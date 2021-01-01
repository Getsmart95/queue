import React from 'react';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import Login from "./pages/Login/Login";
import Tablo from "./pages/Queue/Tablo";
import Manager from "./pages/Main"
import Queue from "./pages/Queue/Store"
// import Operation from "./components/operation"
// import Account from "./components/accounts"
// import Transfer from "./components/transfer"
// import Service from "./components/service"

function App() {
  return (<Router>
    <div className="App">
  
          <Switch>
            <Route exact path='/' component={Login} />
            <Route path="/sign-in" component={Login} />
            <Route path="/tablo" component={Tablo} />
            <Route path="/main" component={Manager} />
            <Route path="/store" component={Queue} />
            {/* <Route path="/account" component={Account} /> */}
            {/* <Route path="/transfer" component={Transfer} /> */}
            {/* <Route path="/service" component={Service} /> */}
          </Switch>
        </div>

    </Router>
  );
}

export default App;