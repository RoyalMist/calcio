import {Form, Formik} from "formik";
import React, {useState} from "react";
import SimpleField from "../../components/simple-field";

const user = {
    name: "Thibault Fouache",
    imageUrl: "/images/icon.png",
};

interface ProfileFormValues {
    name: string;
    password: string;
}

function Profile() {
    const [enableNotifications, setEnableNotifications] = useState(true);
    const initialValues: ProfileFormValues = {
        name: "",
        password: "",
    };

    return (
        <div className="divide-y divide-gray-200">
            <Formik initialValues={initialValues} onSubmit={(values) => {
            }}>
                <Form className="lg:col-span-9">
                    <div className="px-4 py-6 sm:p-6 lg:pb-8">
                        <div>
                            <h2 className="text-lg font-medium leading-6 text-gray-900">
                                Profile
                            </h2>
                        </div>
                        <div className="flex flex-col mt-6 lg:flex-row">
                            <div className="flex-grow space-y-6">
                                <div className="col-span-12 sm:col-span-6">
                                    <SimpleField type="text" name="name" placeholder="John">
                                        Name
                                    </SimpleField>
                                </div>
                            </div>
                        </div>
                        <div className="flex flex-col mt-6 lg:flex-row">
                            <div className="flex-grow space-y-6">
                                <div className="col-span-12 sm:col-span-6">
                                    <SimpleField type="text" name="name" placeholder="John">
                                        Password
                                    </SimpleField>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="flex justify-end pb-4 sm:px-6">
                        <button
                            type="submit"
                            className="w-full mt-2 md:w-1/6 btn btn-primary"
                        >
                            Save
                        </button>
                    </div>
                </Form>
            </Formik>
        </div>
    );
}

export default Profile;
