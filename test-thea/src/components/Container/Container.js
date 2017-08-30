import React, { Component } from 'react';
import '../../App.css';

class Container extends Component{

	render(){
		return(
			<div className="fondo">
				<h2 className="blackColor">{this.props.title}</h2>

				<p className="App-intro">
					{this.props.saludo} {this.props.content}
				</p>
			</div>

		);
	}
}

export default Container;
