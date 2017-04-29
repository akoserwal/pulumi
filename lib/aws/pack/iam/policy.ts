// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as coconut from "@coconut/coconut";

import {Group} from "./group";
import {Role} from "./role";
import {User} from "./user";

export interface InlinePolicy {
    policyDocument: any;
    policyName: string;
}

export class Policy extends coconut.Resource implements PolicyArgs {
    public readonly name: string;
    public policyDocument: any;
    public policyName: string;
    public groups?: Group[];
    public roles?: Role[];
    public users?: User[];

    constructor(name: string, args: PolicyArgs) {
        super();
        if (name === undefined) {
            throw new Error("Missing required resource name");
        }
        this.name = name;
        if (args.policyDocument === undefined) {
            throw new Error("Missing required argument 'policyDocument'");
        }
        this.policyDocument = args.policyDocument;
        if (args.policyName === undefined) {
            throw new Error("Missing required argument 'policyName'");
        }
        this.policyName = args.policyName;
        this.groups = args.groups;
        this.roles = args.roles;
        this.users = args.users;
    }
}

export interface PolicyArgs {
    policyDocument: any;
    policyName: string;
    groups?: Group[];
    roles?: Role[];
    users?: User[];
}


