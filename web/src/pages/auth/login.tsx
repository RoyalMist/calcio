import {Form, Formik} from "formik";
import React, {useEffect} from "react";
import Spinner from "../../components/spinner";
import Split from "../../components/split";
import useAuthStore from "../../hooks/useAuthStore";
import {useMutation} from "react-query";
import {API} from "../../utils/api";
import SimpleField from "../../components/simple-field";

interface LoginFormValues {
    name: string;
    password: string;
}

function Login() {
    const initialValues: LoginFormValues = {name: "", password: ""};
    const authStore = useAuthStore();
    const signIn = useMutation(async (values: LoginFormValues) => {
        const {data} = await API.post("/auth/login", values)
        return data;
    })

    useEffect(() => {
        if (!!signIn.data) {
            authStore.setToken(signIn.data);
        }
    }, [signIn.data, authStore]);

    return (
        <>
            <Spinner loading={signIn.isLoading}/>
            <div className="flex min-h-screen bg-white">
                <div className="flex flex-col justify-center flex-1 px-4 py-12 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
                    <div className="flex flex-col justify-center w-full max-w-sm mx-auto lg:w-96">
                        <img src="/images/logo.webp" alt="Calcio"/>
                        <div className="mt-8">
                            <Split>Login</Split>
                            <div className="mt-6">
                                <Formik
                                    initialValues={initialValues}
                                    onSubmit={(values) => {
                                        signIn.mutate({name: values.name, password: values.password});
                                    }}>
                                    <Form className="space-y-6">
                                        <SimpleField
                                            type="text"
                                            name="name"
                                            placeholder="John"
                                        >
                                            Name
                                        </SimpleField>
                                        <SimpleField
                                            type="password"
                                            name="password"
                                            placeholder="**********"
                                        >
                                            Password
                                        </SimpleField>
                                        <div>
                                            <button type="submit" className="w-full btn btn-primary">
                                                Sign in
                                            </button>
                                        </div>
                                    </Form>
                                </Formik>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="relative flex-1 hidden w-0 lg:block">
                    <img
                        className="absolute inset-0 object-cover w-full h-full"
                        src="/images/landing.webp"
                        alt=""
                    />
                </div>
            </div>
        </>
    );
}

export default Login;