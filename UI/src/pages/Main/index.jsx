import React from 'react';
import {Switch, Route} from 'react-router-dom';
import Admin from "./AdminMain"
import Manager from "./ManagerMain"

export default function Routes() {
    return (
        <Switch>
            {/* <Route exact path="/main" component={Admin}/> */}
            <Route exact path="/main" component={Manager}/>
        </Switch>
    )
}
