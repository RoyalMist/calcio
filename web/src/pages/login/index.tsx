import React, {useEffect} from "react";
import Spinner from "../../components/Spinner";
import Split from "../../components/Split";
import {useMutation} from "react-query";
import landing from "../../../images/landing.webp";
import logo from "../../../images/logo.webp";
import {AuthActionKind, useAuth} from "../../stores/authentication";
import {api_login, AuthenticationService} from "../../gen";
import {SubmitHandler, useForm} from "react-hook-form";
import FormField from "../../components/FormField";

function Login() {
    const {dispatcher} = useAuth();
    const {register, handleSubmit, formState: {errors}} = useForm<api_login>();
    const onSubmit: SubmitHandler<api_login> = data => signIn.mutate(data);

    const signIn = useMutation(async (values: api_login) => {
        return await AuthenticationService.postAuthenticationService(values);
    });

    useEffect(() => {
        if (!!signIn.data) {
            dispatcher({type: AuthActionKind.SET, token: signIn.data})
        }
    }, [signIn.data, dispatcher]);

    return (
        <>
            <Spinner loading={signIn.isLoading}/>
            <div className="flex min-h-screen bg-white">
                <div className="flex flex-col justify-center flex-1 px-4 py-12 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
                    <div className="flex flex-col justify-center w-full max-w-sm mx-auto lg:w-96">
                        <img src={logo} alt="Calcio"/>
                        <div className="mt-8">
                            <Split>Login</Split>
                            <div className="mt-6">
                                <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
                                    <FormField form={register("name", {required: true})} error={errors.name} errorMessage="Please fill a name !" placeholder="John">Name</FormField>
                                    <FormField form={register("password", {required: true})} error={errors.password} errorMessage="Please fill a password !" placeholder="********"
                                               type="password">Password</FormField>
                                    <button type="submit" className="w-full btn btn-primary">Sign in</button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="relative flex-1 hidden w-0 lg:block">
                    <img
                        className="absolute inset-0 object-cover w-full h-full"
                        src={landing}
                        alt=""
                    />
                </div>
            </div>
        </>
    );
}

export default Login;
