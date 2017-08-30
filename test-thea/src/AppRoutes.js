import React from 'react';
import { Route, Switch } from 'react-router-dom';

//Components
import App from './App';
import Home from './components/Home/Home';

const AppRoutes = () =>

	<App>
		<Switch>
			<Route exact path='/' component={Home}/>
			<Route component={Home}/>
		</Switch>
	</App>;

export default AppRoutes;