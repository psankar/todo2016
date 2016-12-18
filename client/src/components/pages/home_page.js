import React, {Component} from 'react';

import NewTask from '../new_task';
import TaskList from '../task_list';

class HomePage extends Component {
    render() {
        return (
            <div>
                HomePage
                <TaskList/>
                <NewTask/>
            </div>
        );

    }
}

export default HomePage;