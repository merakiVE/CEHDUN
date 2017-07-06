import { Component } from '@angular/core';
import { User } from '../models/user';

@Component({
    selector: 'api',
    templateUrl: './api.component.html'
})
export class ApiComponent{

	public title:string;
	public user: User;

	constructor(){
		this.title = 'API';
		this.user = new User("", "", "");
	}

	ngOnInit(){
		console.log("Cargado el componente");
	}

	checkUser(){
		console.log(this.user);
		this.user = new User("", "", "");
	}
}