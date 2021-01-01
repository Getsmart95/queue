import React from 'react';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import Login from "./components/login";
import SignUp from "./components/signup";
import Operation from "./components/operation"
import Account from "./components/accounts"
import Transfer from "./components/transfer"
import Service from "./components/service"

function App() {
  return (<Router>
    <div className="App">
      
      <div className="auth-wrapper">
        <div className="auth-inner">
          <Switch>
            <Route exact path='/' component={Login} />
            <Route path="/sign-in" component={Login} />
            <Route path="/sign-up" component={SignUp} />
            <Route path="/operation" component={Operation} />
            <Route path="/account" component={Account} />
            <Route path="/transfer" component={Transfer} />
            <Route path="/service" component={Service} />
          </Switch>
        </div>
      </div>
    </div></Router>
  );
}

export default App;