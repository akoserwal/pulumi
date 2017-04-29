// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as coconut from "@coconut/coconut";

import {Group} from "./group";
import {InlinePolicy, Policy} from "./policy";

export interface LoginProfile {
    password: string;
    passwordResetRequired?: boolean;
}

export class User extends coconut.Resource implements UserArgs {
    public readonly name: string;
    public readonly userName?: string;
    public groups?: Group[];
    public loginProfile?: LoginProfile;
    public managedPolicies?: Policy[];
    public path?: string;
    public policies?: InlinePolicy[];

    constructor(name: string, args: UserArgs) {
        super();
        if (name === undefined) {
            throw new Error("Missing required resource name");
        }
        this.name = name;
        this.userName = args.userName;
        this.groups = args.groups;
        this.loginProfile = args.loginProfile;
        this.managedPolicies = args.managedPolicies;
        this.path = args.path;
        this.policies = args.policies;
    }
}

export interface UserArgs {
    readonly userName?: string;
    groups?: Group[];
    loginProfile?: LoginProfile;
    managedPolicies?: Policy[];
    path?: string;
    policies?: InlinePolicy[];
}


