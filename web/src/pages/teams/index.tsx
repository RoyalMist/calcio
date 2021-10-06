import React, {useState} from "react";
import SectionHeader from "../../components/section-header";
import {useAuth} from "../../stores/authentication";
import Spinner from "../../components/spinner";
import SlideOver from "../../components/slide-over";
import {ent_Team, TeamsService, UsersService} from "../../gen";
import {useMutation, useQuery, useQueryClient} from "react-query";

const Teams = () => {
    const {token} = useAuth();
    const [open, setOpen] = useState(false);
    const [selectedTeammate, setSelectedTeammate] = useState<string>("");
    const newTeam = () => {
        setSelectedTeammate("");
        setOpen(true);
    }

    const queryClient = useQueryClient();
    const usersQuery = useQuery('users', async () => UsersService.getUsersService(token));
    const teamsQuery = useQuery('teams', async () => TeamsService.getTeamsService(token))
    const createTeam = useMutation('teams', async (teammateId?: string) => {
        return await TeamsService.putTeamsService(token, teammateId)
    }, {
        onSuccess: async () => {
            close();
            return await queryClient.invalidateQueries('teams');
        }
    });

    return (
        <>
            <Spinner loading={usersQuery.isLoading || teamsQuery.isLoading || createTeam.isLoading}/>
            <SectionHeader action={newTeam}>Teams</SectionHeader>
            <SlideOver open={open} close={() => setOpen(false)} title="Create New Team">
                TODO
                {/*<Formik
                    initialValues={selectedUser}
                    onSubmit={(values) => {
                        saveUser.mutate({...values});
                    }}>
                    <Form className="space-y-6">
                        {selectedUser.id === undefined &&
                        <SimpleField
                            type="text"
                            name="name"
                            placeholder="John"
                        >
                            Name
                        </SimpleField>
                        }
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
                </Formik>*/}
            </SlideOver>
            <ul role="list" className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {teamsQuery.isSuccess && teamsQuery.data.map((team: ent_Team) => (
                    <li
                        key={team.id}
                        className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow divide-y divide-gray-200"
                    >
                        <div className="flex-1 flex flex-col p-8">
                                <span className="inline-flex items-center justify-center h-14 w-14 rounded-full bg-gray-500">
                                    <span className="text-xl font-medium leading-none text-white">{team.name?.substr(0, 2).toUpperCase()}</span>
                                </span>
                            <h3 className="mt-6 text-gray-900 text-sm font-medium">{team.name}</h3>
                        </div>
                    </li>
                ))}
            </ul>
        </>

    );
};

export default Teams;
