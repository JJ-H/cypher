export namespace models {
	
	export class Credential {
	    id: number;
	    domain: string;
	    username: string;
	    password: string;
	    note: string;
	
	    static createFrom(source: any = {}) {
	        return new Credential(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.domain = source["domain"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.note = source["note"];
	    }
	}

}

